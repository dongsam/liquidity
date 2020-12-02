package keeper
import (
	"github.com/tendermint/liquidity/x/liquidity/types"
	"context"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the distribution MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}


//func (k msgServer) CreateLiquidityPool(context.Context, *types.MsgCreateLiquidityPoolRequest) (*types.MsgCreateLiquidityPoolResponse, error){}
func (k msgServer) CreateLiquidityPool(goCtx context.Context, msg *types.MsgCreateLiquidityPoolRequest) (*types.MsgCreateLiquidityPoolResponse, error) {
	return nil, nil
}
func (k msgServer) DepositToLiquidityPool(goCtx context.Context, msg *types.MsgDepositToLiquidityPoolRequest) (*types.MsgDepositToLiquidityPoolResponse, error) {
	return nil, nil
}
func (k msgServer) WithdrawFromLiquidityPool(goCtx context.Context, msg *types.MsgWithdrawFromLiquidityPoolRequest) (*types.MsgWithdrawFromLiquidityPoolResponse, error) {
	return nil, nil
}
func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwapRequest) (*types.MsgSwapResponse, error) {
	return nil, nil
}
