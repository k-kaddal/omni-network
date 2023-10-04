package keeper

import (
	"github.com/k-kaddal/omni/x/ethbridge/types"
)

var _ types.QueryServer = Keeper{}
