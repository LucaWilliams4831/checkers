package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/LucaWilliams4831/checkers/testutil/keeper"
	"github.com/LucaWilliams4831/checkers/testutil/nullify"
	"github.com/LucaWilliams4831/checkers/x/checkers/keeper"
	"github.com/LucaWilliams4831/checkers/x/checkers/types"
)

func createTestNextGame(keeper *keeper.Keeper, ctx sdk.Context) types.NextGame {
	item := types.NextGame{}
	keeper.SetNextGame(ctx, item)
	return item
}

func TestNextGameGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	item := createTestNextGame(keeper, ctx)
	rst, found := keeper.GetNextGame(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNextGameRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	createTestNextGame(keeper, ctx)
	keeper.RemoveNextGame(ctx)
	_, found := keeper.GetNextGame(ctx)
	require.False(t, found)
}
