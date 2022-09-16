package blog1_test

import (
	"testing"

	keepertest "blog1/testutil/keeper"
	"blog1/testutil/nullify"
	"blog1/x/blog1"
	"blog1/x/blog1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Blog1Keeper(t)
	blog1.InitGenesis(ctx, *k, genesisState)
	got := blog1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
