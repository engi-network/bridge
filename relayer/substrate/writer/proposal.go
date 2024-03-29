package writer

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

type SubstrateProposal struct {
	DepositNonce types.U64
	Call         types.Call
	SourceId     types.U8
	ResourceId   types.Bytes32
	Method       string
}

// encode takes only nonce and call and encodes them for storage queries
func (p *SubstrateProposal) Encode() ([]byte, error) {
	return codec.Encode(struct {
		types.U64
		types.Call
	}{p.DepositNonce, p.Call})
}
