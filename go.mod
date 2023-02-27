module github.com/engi-network/bridge

go 1.15

replace github.com/ChainSafe/chainbridge-core => github.com/engi-network/sygma-core v0.0.8

//
// replace github.com/ChainSafe/chainbridge-celo-module => ../chainbridge-celo-module

require (
	github.com/ChainSafe/chainbridge-celo-module v0.0.0-20220121131741-69b2ecf7dec5
	github.com/ChainSafe/chainbridge-core v0.0.0-20220120162654-c03a4d159125
	github.com/ChainSafe/chainbridge-utils v1.0.6
	github.com/ChainSafe/log15 v1.0.0
	github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.12
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.10.1
	github.com/status-im/keycard-go v0.0.0-20211004132608-c32310e39b86
	github.com/stretchr/testify v1.7.0
)
