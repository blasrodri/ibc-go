package wasm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
)

func (cs ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore,
	substituteClientStore sdk.KVStore, substituteClient exported.ClientState,
) error {
	return nil
}
