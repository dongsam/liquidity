package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/liquidity/x/liquidity/types"
)

func (k Keeper) TrackPoolBatch(ctx sdk.Context) {
	k.IterateAllLiquidityPoolBatches(ctx, func(liquidityPoolBatch types.LiquidityPoolBatch) bool {
		if liquidityPoolBatch.ExecutionStatus {
			k.
		}
		// TODO: check
		// TODO: delete
		liquidityPoolBatch.BeginHeight
		return false
	})

	// TODO: init?
}