package keeper_test

import (
	"testing"

	testkeeper "github.com/k-kaddal/omni/testutil/keeper"
	"github.com/k-kaddal/omni/x/omni/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OmniKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
