package keeper

import (
	"testchain/x/faucet/types"
)

var _ types.QueryServer = Keeper{}
