package keeper

import (
	"github.com/k-kaddal/omni/x/ethereumbridge/types"
)

var _ types.QueryServer = Keeper{}
