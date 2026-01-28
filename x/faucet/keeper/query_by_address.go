package keeper

import (
	"context"

	"testchain/x/faucet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ByAddress(goCtx context.Context, req *types.QueryByAddressRequest) (*types.QueryByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse the address
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	// Get all requests for this address
	requestsData := k.GetRequestsByAddress(ctx, addr)

	// Convert RequestData to proto Request
	var requests []types.Request
	for _, req := range requestsData {
		requests = append(requests, types.Request{
			Amount: req.Amount,
			Height: req.Height,
		})
	}

	// Get total requested
	total := k.GetTotalRequested(ctx, addr)

	return &types.QueryByAddressResponse{
		Requests: requests,
		Total:    total,
	}, nil
}
