package keeper_test

import (
	"context"
	"testing"

	keepertest "blog1/testutil/keeper"
	"blog1/x/blog1/keeper"
	"blog1/x/blog1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Blog1Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
