// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package listener

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type EventExampleRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.U32
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventRadClaimsClaimed is emitted when RAD Tokens have been claimed
type EventRadClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventRadClaimsRootHashStored is emitted when RootHash has been stored for the correspondent RAD Claims batch
type EventRadClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

type Erc721Token struct {
	Id       types.U256
	Metadata types.Bytes
}

type RegistryId types.H160
type TokenId types.U256

type AssetId struct {
	RegistryId RegistryId
	TokenId    TokenId
}

// EventNftTransferred is emitted when the ownership of the asset has been transferred to the account
type EventNftTransferred struct {
	Phase      types.Phase
	RegistryId RegistryId
	AssetId    AssetId
	Who        types.AccountID
	Topics     []types.Hash
}

// EventRegistryMint is emitted when successfully minting an NFT
type EventRegistryMint struct {
	Phase      types.Phase
	RegistryId RegistryId
	TokenId    TokenId
	Topics     []types.Hash
}

// EventRegistryRegistryCreated is emitted when successfully creating a NFT registry
type EventRegistryRegistryCreated struct {
	Phase      types.Phase
	RegistryId RegistryId
	Topics     []types.Hash
}

// EventRegistryTmp is emitted only for testing
type EventRegistryTmp struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Amount       types.U256
	Recipient    types.Bytes
	Topics       []types.Hash
}

type EventNonFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	TokenId      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventGenericTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventRelayerThresholdChanged struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventChainWhitelisted struct {
	Phase   types.Phase
	ChainId types.U8
	Topics  []types.Hash
}

type EventRelayerAdded struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventRelayerRemoved struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventVoteFor struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventVoteAgainst struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventProposalApproved struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalRejected struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalSucceeded struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalFailed struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventJobCreated struct {
	Phase        types.Phase
    JobId        types.U64
    AccountId    types.AccountID
	Topics       []types.Hash
}

type EventJobIdGenerated struct {
	Phase        types.Phase
    JobId        types.U64
	Topics       []types.Hash
}

type Events struct {
	types.EventRecords
	ChainBridge_FungibleTransfer        []EventFungibleTransfer               //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventNonFungibleTransfer            //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventGenericTransfer                //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventRelayerThresholdChanged        //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainWhitelisted               //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventRelayerAdded                   //nolint:stylecheck,golint
	ChainBridge_RelayerRemoved          []EventRelayerRemoved                 //nolint:stylecheck,golint
	ChainBridge_VoteFor                 []EventVoteFor                        //nolint:stylecheck,golint
	ChainBridge_VoteAgainst             []EventVoteAgainst                    //nolint:stylecheck,golint
	ChainBridge_ProposalApproved        []EventProposalApproved               //nolint:stylecheck,golint
	ChainBridge_ProposalRejected        []EventProposalRejected               //nolint:stylecheck,golint
	ChainBridge_ProposalSucceeded       []EventProposalSucceeded              //nolint:stylecheck,golint
	ChainBridge_ProposalFailed          []EventProposalFailed                 //nolint:stylecheck,golint
    Jobs_JobCreated                     []EventJobCreated                     //nolint:stylecheck,golint
    Jobs_JobIdGenerated                 []EventJobIdGenerated                 //nolint:stylecheck,golint
	Example_Remark                      []EventExampleRemark                  //nolint:stylecheck,golint
	Nfts_DepositAsset                   []EventNFTDeposited                   //nolint:stylecheck,golint
	Council_Proposed                    []types.EventCouncilProposed       //nolint:stylecheck,golint
	Council_Voted                       []types.EventCouncilVoted          //nolint:stylecheck,golint
	Council_Approved                    []types.EventCouncilApproved       //nolint:stylecheck,golint
	Council_Disapproved                 []types.EventCouncilDisapproved    //nolint:stylecheck,golint
	Council_Executed                    []types.EventCouncilExecuted       //nolint:stylecheck,golint
	Council_MemberExecuted              []types.EventCouncilMemberExecuted //nolint:stylecheck,golint
	Council_Closed                      []types.EventCouncilClosed         //nolint:stylecheck,golint
	Fees_FeeChanged                     []EventFeeChanged                     //nolint:stylecheck,golint
	MultiAccount_NewMultiAccount        []EventNewMultiAccount                //nolint:stylecheck,golint
	MultiAccount_MultiAccountUpdated    []EventMultiAccountUpdated            //nolint:stylecheck,golint
	MultiAccount_MultiAccountRemoved    []EventMultiAccountRemoved            //nolint:stylecheck,golint
	MultiAccount_NewMultisig            []EventNewMultisig                    //nolint:stylecheck,golint
	MultiAccount_MultisigApproval       []EventMultisigApproval               //nolint:stylecheck,golint
	MultiAccount_MultisigExecuted       []EventMultisigExecuted               //nolint:stylecheck,golint
	MultiAccount_MultisigCancelled      []EventMultisigCancelled              //nolint:stylecheck,golint
	TreasuryReward_TreasuryMinting      []EventTreasuryMinting                //nolint:stylecheck,golint
	Nft_Transferred                     []EventNftTransferred                 //nolint:stylecheck,golint
	RadClaims_Claimed                   []EventRadClaimsClaimed               //nolint:stylecheck,golint
	RadClaims_RootHashStored            []EventRadClaimsRootHashStored        //nolint:stylecheck,golint
	Registry_Mint                       []EventRegistryMint                   //nolint:stylecheck,golint
	Registry_RegistryCreated            []EventRegistryRegistryCreated        //nolint:stylecheck,golint
	Registry_RegistryTmp                []EventRegistryTmp                    //nolint:stylecheck,golint
}
