package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/tendermint/liquidity/types"
)

func registerTxRoutes(clientCtx client.Context, r *mux.Router) {
	// create liquidityPool
	r.HandleFunc(fmt.Sprintf("/liquidity/pool"), newLiquidityPoolHandlerFn(clientCtx)).Methods("POST")
	// deposit to liquidityPool
	r.HandleFunc(fmt.Sprintf("/liquidity/pool/{%s}/deposit", RestPoolID), newDepositLiquidityPoolHandlerFn(clientCtx)).Methods("POST")
	// withdraw from liquidityPool
	r.HandleFunc(fmt.Sprintf("/liquidity/pool/{%s}/withdraw", RestPoolID), newWithdrawLiquidityPoolHandlerFn(clientCtx)).Methods("POST")
}

// TODO: add detailed logic to each handler

func newLiquidityPoolHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateLiquidityPoolReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		poolCreator, e := sdk.AccAddressFromBech32(req.PoolCreator)
		if e != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, e.Error())
			return
		}

		depositToken, ok := sdk.NewIntFromString(req.DepositTokensAmount)
		if !ok || !depositToken.IsPositive() {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "token amount: "+req.DepositTokensAmount)
			return
		}

		msg := types.NewMsgCreateLiquidityPool()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

// HTTP request handler to add liquidity.
func newDepositLiquidityPoolHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		poolID := vars[RestPoolID]

		var req DepositLiquidityPoolReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgDepositToLiquidityPool()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

// HTTP request handler to remove liquidity.
func newWithdrawLiquidityPoolHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		poolID := vars[RestPoolID]

		var req WithdrawLiquidityPoolReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgWithdrawFromLiquidityPool()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

// WithdrawLiquidityPoolReq defines the properties of a Deposit from liquidity Pool request's body
type CreateLiquidityPoolReq struct {
	BaseReq             rest.BaseReq `json:"base_req" yaml:"base_req"`
	PoolCreator         string       `json:"pool_creator" yaml:"pool_creator"`
	PoolTypeIndex       string       `json:"pool_type_index" yaml:"pool_type_index"`
	ReserveTokenDenoms  string       `json:"reserve_token_denoms" yaml:"reserve_token_denoms"`
	DepositTokensAmount string       `json:"deposit_token_amount" yaml:"deposit_token_amount"`
}

// WithdrawLiquidityPoolReq defines the properties of a Deposit from liquidity Pool request's body
type WithdrawLiquidityPoolReq struct {
	BaseReq         rest.BaseReq `json:"base_req" yaml:"base_req"`
	Withdrawer      string       `json:"withdrawer" yaml:"withdrawer"`
	PoolID          string       `json:"pool_id" yaml:"pool_id"`
	PoolTokenAmount string       `json:"pool_token_amount" yaml:"pool_token_amount"`
}

// DepositLiquidityPoolReq defines the properties of a Deposit liquidity request's body
type DepositLiquidityPoolReq struct {
	BaseReq             rest.BaseReq `json:"base_req" yaml:"base_req"`
	Depositor           string       `json:"depositor" yaml:"depositor"`
	PoolID              string       `json:"pool_id" yaml:"pool_id"`
	DepositTokensAmount string       `json:"deposit_tokens_amount" yaml:"deposit_tokens_amount"`
}

// DepositLiquidityPoolReq defines the properties of a Deposit liquidity request's body
type SwapReq struct {
	BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
	SwapRequester string       `json:"swap_requester" yaml:"swap_requester"`
	PoolID        string       `json:"pool_id" yaml:"pool_id"`
	PoolTypeIndex string       `json:"pool_type_index" yaml:"pool_type_index"`
	SwapType      string       `json:"swap_type" yaml:"swap_type"`
	OfferToken    string       `json:"offer_token" yaml:"offer_token"`
	DemandToken   string       `json:"demand_token" yaml:"demand_token"`
	OrderPrice    string       `json:"order_price" yaml:"order_price"`
}
