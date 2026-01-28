package keeper

import (
	"context"

	"testchain/x/faucet/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Request(goCtx context.Context, msg *types.MsgRequest) (*types.MsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("message is nil")
	}

	if msg.Amount == 0 {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("amount must be greater than zero")
	}

	params := k.Keeper.GetParams(ctx)

	// check per pre-request limit
	if params.MaxPerRequest > 0 && msg.Amount > params.MaxPerRequest {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("amount exceeds max per request limit of %d", params.MaxPerRequest)
	}

	// resolve signer address
	var requester sdk.AccAddress

	if msg.Creator != "" {
		addr, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, sdkerrors.ErrInvalidAddress.Wrap("invalid creator address")
		}
		requester = addr
	} else {
		return nil, sdkerrors.ErrInvalidAddress.Wrap("creator address is required")
	}

	// check against per-address limit
	alreadyRequested := k.Keeper.GetTotalRequested(ctx, requester)

	if params.MaxPerAddress > 0 && (alreadyRequested+msg.Amount) > params.MaxPerAddress {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("request exceeds max per address limit of %d", params.MaxPerAddress)
	}

	// transfer from module to requester
	coin := sdk.NewCoin(params.DefaultDenom, math.NewIntFromUint64(msg.Amount))
	coins := sdk.NewCoins(coin)

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requester, coins); err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "failed to send coins from module to account: %v", err)
	}

	// update total requested
	if err := k.Keeper.AppendRequest(ctx, requester, msg.Amount, ctx.BlockHeight()); err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "failed to append request: %v", err)
	}

	if err := k.Keeper.AddToTotalRequested(ctx, requester, msg.Amount); err != nil {
		_ = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, coins)
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "failed to update total requested: %v", err)
	}

	return &types.MsgRequestResponse{}, nil
}
