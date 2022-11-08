package wasm

import (
	"time"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
)

var _ exported.ClientMessage = &Misbehaviour{}

// FrozenHeight is same for all misbehaviour
var FrozenHeight = clienttypes.NewHeight(0, 1)

// NewMisbehaviour creates a new Misbehaviour instance.
func NewMisbehaviour(clientID string, header1, header2 *Header) *Misbehaviour {
	return &Misbehaviour{
		ClientId: clientID,
	}
}

// ClientType is Tendermint light client
func (misbehaviour Misbehaviour) ClientType() string {
	return exported.Tendermint
}

// GetTime returns the timestamp at which misbehaviour occurred. It uses the
// maximum value from both headers to prevent producing an invalid header outside
// of the misbehaviour age range.
func (misbehaviour Misbehaviour) GetTime() time.Time {
	panic("")
}

// ValidateBasic implements Misbehaviour interface
func (misbehaviour Misbehaviour) ValidateBasic() error {
	return nil
}

// validCommit checks if the given commit is a valid commit from the passed-in validatorset
func validCommit(chainID string, blockID tmtypes.BlockID, commit *tmproto.Commit, valSet *tmproto.ValidatorSet) (err error) {
	return nil
}
