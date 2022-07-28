// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
"github.com/engi-network/bridge/relayer/substrate/writer"
)

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = writer.BridgePalletName + ".add_relayer"
var SetResourceMethod Method = writer.BridgePalletName + ".set_resource"
var SetThresholdMethod Method = writer.BridgePalletName + ".set_threshold"
var WhitelistChainMethod Method = writer.BridgePalletName + ".whitelist_chain"
var ExampleTransferNativeMethod Method = "Example.transfer_native"
var ExampleTransferErc721Method Method = "Example.transfer_erc721"
var ExampleTransferHashMethod Method = "Example.transfer_hash"
var ExampleMintErc721Method Method = "Example.mint_erc721"
var ExampleTransferMethod Method = "Example.transfer"
var ExampleRemarkMethod Method = "Example.remark"
var Erc721MintMethod Method = "Erc721.mint"
var SudoMethod Method = "Sudo.sudo"
