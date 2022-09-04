package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/LucaWilliams4831/checkers/testutil/keeper"
	"github.com/LucaWilliams4831/checkers/testutil/nullify"
	"github.com/LucaWilliams4831/checkers/x/checkers/types"
)

func TestNextGameQuery(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNextGame(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNextGameRequest
		response *types.QueryGetNextGameResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNextGameRequest{},
			response: &types.QueryGetNextGameResponse{NextGame: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NextGame(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
