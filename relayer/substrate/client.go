// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
        "errors"
	"fmt"
	"math/big"
        "sync"

	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/chainbridge-core/types"
	"github.com/engi-network/bridge/relayer/substrate/writer"
	"github.com/ChainSafe/log15"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	grpcTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/author"
)

var TerminatedError = errors.New("terminated")

// Client is a container for all the components required to submit extrinsics
// TODO: Perhaps this would benefit an interface so we can interchange Connection and a client like this
type Client struct {
	Api     *gsrpc.SubstrateAPI
	Meta    *grpcTypes.Metadata
        Metalock sync.RWMutex
	Genesis grpcTypes.Hash
	Key     *signature.KeyringPair
        Nonce   grpcTypes.U32
        Noncelock sync.Mutex
}

func (c *Client) getLatestNonce() (grpcTypes.U32, error) {
    var acct grpcTypes.AccountInfo
    meta := c.getMetadata()

    key, err := grpcTypes.CreateStorageKey(meta, "System", "Account", c.Key.PublicKey, nil)
    if err != nil {
        return 0, err
    }

    exists, err := c.Api.RPC.State.GetStorageLatest(key, &acct)
    if !exists {
        return 0, fmt.Errorf("account %x not found on chain", c.Key.PublicKey)
    }

    return acct.Nonce, nil
}

func (c *Client) getMetadata() (*grpcTypes.Metadata) {
    c.Metalock.RLock()
    meta := c.Meta
    c.Metalock.RUnlock()
    return meta
}

// Implement the Voter interface
func (c *Client) SubmitTx(method string, args ...interface{}) error {
    // AcknowledgeProposal, prop.DepositNonce, prop.SourceId, prop.ResourceId, prop.Call
    fmt.Println("Submitting call to substrate")

    meta := c.getMetadata()

    call, err := grpcTypes.NewCall(meta, method, args...,)
    if err != nil {
        return fmt.Errorf("failed to construct call: %w", err)
    }

    ext := grpcTypes.NewExtrinsic(call)

    rv, err := c.Api.RPC.State.GetRuntimeVersionLatest()
    if err != nil {
        return err
    }

    c.Noncelock.Lock()
    latest, err := c.getLatestNonce()
    if err != nil {
        c.Noncelock.Unlock()
        return err
    }
    if latest > c.Nonce {
        c.Nonce = latest
    }

    o := grpcTypes.SignatureOptions{
	Era: grpcTypes.ExtrinsicEra{IsMortalEra: false},
	Nonce: grpcTypes.NewUCompactFromUInt(uint64(c.Nonce)),
	Tip: grpcTypes.NewUCompactFromUInt(0),
	SpecVersion: rv.SpecVersion,
	GenesisHash: c.Genesis,
	BlockHash: c.Genesis,
	TransactionVersion: rv.TransactionVersion,

    }

    err = ext.Sign(*c.Key, o)
    if err != nil {
        c.Noncelock.Unlock()
        return err
    }
    sub, err := c.Api.RPC.Author.SubmitAndWatchExtrinsic(ext)
    c.Nonce++
    c.Noncelock.Unlock()

    if err != nil {
        return fmt.Errorf("submission of extrinsic failed %w", err)
    }
    defer sub.Unsubscribe()

    return c.watchSubmission(sub)
}

func (c *Client) watchSubmission(sub *author.ExtrinsicStatusSubscription) error {
    for {
        select {
        case status := <- sub.Chan():
            switch {
            case status.IsInBlock:
		log15.Info("Extrinsic included in block: ", status.AsInBlock.Hex())
                return nil
            case status.IsRetracted:
                return fmt.Errorf("extrinsic retracted %s", status.AsRetracted.Hex())
            case status.IsDropped:
                return fmt.Errorf("extrinsic dropped from network")
            case status.IsInvalid:
                return fmt.Errorf("extrinsic invalid")
            default:
		fmt.Println("Waiting on extrinsic status")
            }
        case err := <-sub.Err():
            return fmt.Errorf("Extrinsic subscription error: ", err)
       }
    }
}

func (c *Client) GetVoterAccountID() (*grpcTypes.AccountID, error) {
    var acctId, err = grpcTypes.NewAccountID(c.Key.PublicKey)

    if err != nil {
        return nil, err
    }

    return acctId, nil
}

func (c *Client) GetMetadata() (meta grpcTypes.Metadata) {
    meta = *c.getMetadata()
    return meta
}

func (c *Client) ResolveResourceId(resourceId types.ResourceID) (string, error) {
    var res []byte
    meta := c.getMetadata()

    key, err := grpcTypes.CreateStorageKey(meta, writer.BridgeStoragePrefix, "Resources", resourceId[:], nil)
    if err != nil {
        return "", err
    }

    exists, err := c.Api.RPC.State.GetStorageLatest(key, &res)
    if !exists {
        return "", fmt.Errorf("resource %x not found on chain", resourceId)
    }

    return string(res), nil
}

func (c *Client) GetProposalStatus(sourceID, proposalBytes []byte) (bool, *writer.VoteState, error) {
    var voteRes writer.VoteState
    meta := c.getMetadata()
    key, err := grpcTypes.CreateStorageKey(meta, writer.BridgeStoragePrefix, "Votes", sourceID[:], proposalBytes)
    if err != nil {
        return false, nil, err
    }

    exists, err := c.Api.RPC.State.GetStorageLatest(key, &voteRes)
    if !exists {
        return false, nil, nil
    }

    return true, &voteRes, nil
}


// Implement the SubstrateReader interface
func (c *Client) GetHeaderLatest() (*grpcTypes.Header, error) {
	return c.Api.RPC.Chain.GetHeaderLatest()
}

func (c *Client) GetBlockHash(blockNumber uint64) (grpcTypes.Hash, error) {
    return c.Api.RPC.Chain.GetBlockHash(blockNumber)
}

func (c *Client) GetBlockEvents(hash grpcTypes.Hash, target interface{}) error {
    meta := c.getMetadata()
    key, err := grpcTypes.CreateStorageKey(meta, "System", "Events", nil, nil)
    if err != nil {
        return err
    }
    var records grpcTypes.EventRecordsRaw
    _, err = c.Api.RPC.State.GetStorage(key, &records, hash)
    if err != nil {
        return err
    }
    err = records.DecodeEventRecords(meta, target)
    if err != nil {
        return err
    }
    return nil
}

func (c *Client) UpdateMetatdata() error {
    meta, err := c.Api.RPC.State.GetMetadataLatest()
    if err != nil {
        return err
    }
    c.Metalock.Lock()
    c.Meta = meta
    c.Metalock.Unlock()
    return nil
}

func CreateClient(key *signature.KeyringPair, endpoint string) (*Client, error) {
	c := &Client{Key: key}
	api, err := gsrpc.NewSubstrateAPI(endpoint)
	if err != nil {
		return nil, err
	}
	c.Api = api

	// Fetch metadata
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	c.Meta = meta

	// Fetch genesis hash
	genesisHash, err := c.Api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}
	c.Genesis = genesisHash

	return c, nil
}

// Admin calls

func (c *Client) SetRelayerThreshold(threshold grpcTypes.U32) error {
	log15.Info("Setting threshold", "threshold", threshold)
	return SubmitSudoTx(c, SetThresholdMethod, threshold)
}

func (c *Client) AddRelayer(relayer grpcTypes.AccountID) error {
	log15.Info("Adding relayer", "accountId", relayer)
	return SubmitSudoTx(c, AddRelayerMethod, relayer)
}

func (c *Client) WhitelistChain(id msg.ChainId) error {
	log15.Info("Whitelisting chain", "chainId", id)
	return SubmitSudoTx(c, WhitelistChainMethod, grpcTypes.U8(id))
}

func (c *Client) RegisterResource(id msg.ResourceId, method string) error {
	log15.Info("Registering resource", "rId", id, "method", []byte(method))
	return SubmitSudoTx(c, SetResourceMethod, grpcTypes.NewBytes32(id), []byte(method))
}

// Standard transfer calls

func (c *Client) InitiateNativeTransfer(amount grpcTypes.U128, recipient []byte, destId msg.ChainId) error {
	log15.Info("Initiating Substrate native transfer", "amount", amount, "recipient", fmt.Sprintf("%x", recipient), "destId", destId)
	return SubmitTx(c, ExampleTransferNativeMethod, amount, recipient, grpcTypes.U8(destId))
}

func (c *Client) InitiateNonFungibleTransfer(tokenId grpcTypes.U256, recipient []byte, destId msg.ChainId) error {
	log15.Info("Initiating Substrate nft transfer", "tokenId", tokenId, "recipient", recipient, "destId", destId)
	return SubmitTx(c, ExampleTransferErc721Method, recipient, tokenId, grpcTypes.U8(destId))
}

func (c *Client) InitiateHashTransfer(hash grpcTypes.Hash, destId msg.ChainId) error {
	log15.Info("Initiating hash transfer", "hash", hash.Hex())
	return SubmitTx(c, ExampleTransferHashMethod, hash, grpcTypes.U8(destId))
}

// Call creation methods for batching

func (c *Client) NewSudoCall(call grpcTypes.Call) (grpcTypes.Call, error) {
	return grpcTypes.NewCall(c.Meta, string(SudoMethod), call)
}

func (c *Client) NewSetRelayerThresholdCall(threshold grpcTypes.U32) (grpcTypes.Call, error) {
	call, err := grpcTypes.NewCall(c.Meta, string(SetThresholdMethod), threshold)
	if err != nil {
		return grpcTypes.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewAddRelayerCall(relayer grpcTypes.AccountID) (grpcTypes.Call, error) {
	call, err := grpcTypes.NewCall(c.Meta, string(AddRelayerMethod), relayer)
	if err != nil {
		return grpcTypes.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewWhitelistChainCall(id msg.ChainId) (grpcTypes.Call, error) {
	call, err := grpcTypes.NewCall(c.Meta, string(WhitelistChainMethod), id)
	if err != nil {
		return grpcTypes.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewRegisterResourceCall(id msg.ResourceId, method string) (grpcTypes.Call, error) {
	call, err := grpcTypes.NewCall(c.Meta, string(SetResourceMethod), id, method)
	if err != nil {
		return grpcTypes.Call{}, err
	}
	return c.NewSudoCall(call)
}

func (c *Client) NewNativeTransferCall(amount grpcTypes.U128, recipient []byte, destId msg.ChainId) (grpcTypes.Call, error) {
	return grpcTypes.NewCall(c.Meta, string(ExampleTransferNativeMethod), amount, recipient, grpcTypes.U8(destId))
}

// Utility methods

func (c *Client) LatestBlock() (uint64, error) {
	head, err := c.Api.RPC.Chain.GetHeaderLatest()
	if err != nil {
		return 0, err
	}
	return uint64(head.Number), nil
}

func (c *Client) MintErc721(tokenId *big.Int, metadata []byte, recipient *signature.KeyringPair) error {
	fmt.Printf("Mint info: account %x amount: %x meta: %x\n", recipient.PublicKey, grpcTypes.NewU256(*tokenId), grpcTypes.Bytes(metadata))

    var acctId, err = grpcTypes.NewAccountID(recipient.PublicKey)

    if err != nil {
        return err
    }

	return SubmitSudoTx(c, Erc721MintMethod, acctId, grpcTypes.NewU256(*tokenId), grpcTypes.Bytes(metadata))
}

func (c *Client) OwnerOf(tokenId *big.Int) (grpcTypes.AccountID, error) {
	var owner grpcTypes.AccountID
	tokenIdBz, err := codec.Encode(grpcTypes.NewU256(*tokenId))
	if err != nil {
		return grpcTypes.AccountID{}, err
	}

	exists, err := QueryStorage(c, "TokenStorage", "TokenOwner", tokenIdBz, nil, &owner)
	if err != nil {
		return grpcTypes.AccountID{}, err
	}
	if !exists {
		return grpcTypes.AccountID{}, fmt.Errorf("token %s doesn't have an owner", tokenId.String())
	}
	return owner, nil
}

func (c *Client) GetDepositNonce(chain msg.ChainId) (uint64, error) {
	var count grpcTypes.U64
	chainId, err := codec.Encode(grpcTypes.U8(chain))
	if err != nil {
		return 0, err
	}
	exists, err := QueryStorage(c, writer.BridgeStoragePrefix, "ChainNonces", chainId, nil, &count)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, nil
	}
	return uint64(count), nil
}
