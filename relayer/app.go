// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package relayer

import (
        "fmt"
	"os"
	"os/signal"
	"syscall"
        "math/big"

	"github.com/ChainSafe/chainbridge-core/chains/evm"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/evmtransaction"
	subclient "github.com/engi-network/bridge/relayer/substrate"
	sublistener "github.com/engi-network/bridge/relayer/substrate/listener"
	subwriter "github.com/engi-network/bridge/relayer/substrate/writer"
	"github.com/ChainSafe/chainbridge-core/chains/substrate"
	grpcTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/ChainSafe/chainbridge-core/config"
	"github.com/ChainSafe/chainbridge-core/keystore"
	"github.com/ChainSafe/chainbridge-core/config/chain"
	"github.com/ChainSafe/chainbridge-core/flags"
	"github.com/ChainSafe/chainbridge-core/lvldb"
	"github.com/ChainSafe/chainbridge-core/opentelemetry"
	"github.com/ChainSafe/chainbridge-core/relayer"
	"github.com/ChainSafe/chainbridge-core/relayer/message"
	"github.com/ChainSafe/chainbridge-core/store"
	"github.com/ChainSafe/chainbridge-core/crypto/sr25519"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func handleFungibleTransfer(msg *message.Message) ([]interface{}) {
    m := make([]interface{}, 0)

    amt := big.NewInt(0).SetBytes(msg.Payload[0].([]byte))
    amount := grpcTypes.NewU128(*amt)
    recipient := grpcTypes.NewAccountID(msg.Payload[1].([]byte))

    m = append(m, recipient)
    m = append(m, amount)
    return m
}

func Run() error {
	errChn := make(chan error)
	stopChn := make(chan struct{})

	configuration, err := config.GetConfig(viper.GetString(flags.ConfigFlagName))
	if err != nil {
		panic(err)
	}
	db, err := lvldb.NewLvlDB(viper.GetString(flags.BlockstoreFlagName))
	if err != nil {
		panic(err)
	}
	blockstore := store.NewBlockStore(db)
fmt.Println(configuration)

	chains := []relayer.RelayedChain{}
	for _, chainConfig := range configuration.ChainConfigs {
		switch chainConfig["type"] {
		case "evm":
			{
				chain, err := evm.SetupDefaultEVMChain(chainConfig, evmtransaction.NewTransaction, blockstore)
				if err != nil {
					panic(err)
				}

				chains = append(chains, chain)
			}
                case "substrate":
                        {
                               config, err := chain.NewSubstrateConfig(chainConfig)
                               if err != nil {
                                       panic(err)
                               }
                               keypair, err := keystore.KeypairFromAddress(config.GeneralChainConfig.From, keystore.SubChain, config.GeneralChainConfig.KeystorePath, config.GeneralChainConfig.Insecure)
                               if err != nil {
                                       panic(err)
                               }
                               krp := keypair.(*sr25519.Keypair).AsKeyringPair()
                               client,err := subclient.CreateClient((*signature.KeyringPair)(krp), config.GeneralChainConfig.Endpoint)
                               if err != nil {
                                       panic(err)
                               }
fmt.Println(client)
                               subListener := sublistener.NewSubstrateListener(client)
                               subListener.RegisterSubscription(message.FungibleTransfer, sublistener.FungibleTransferHandler)
                               writer := subwriter.NewSubstrateWriter(*config.GeneralChainConfig.Id, client)
                               writer.RegisterHandler(message.FungibleTransfer, handleFungibleTransfer)
                               chains = append(chains, substrate.NewSubstrateChain(subListener, writer, blockstore, *config.GeneralChainConfig.Id, config))
                        }
		}
	}

	r := relayer.NewRelayer(chains, &opentelemetry.ConsoleTelemetry{})
	go r.Start(stopChn, errChn)

	sysErr := make(chan os.Signal, 1)
	signal.Notify(sysErr,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	select {
	case err := <-errChn:
		log.Error().Err(err).Msg("failed to listen and serve")
		close(stopChn)
		return err
	case sig := <-sysErr:
		log.Info().Msgf("terminating got [%v] signal", sig)
		return nil
	}
}
