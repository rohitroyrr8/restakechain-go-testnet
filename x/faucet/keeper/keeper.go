package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"testchain/x/faucet/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger
		bankKeeper   types.BankKeeper

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	bankKeeper types.BankKeeper,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		bankKeeper:   bankKeeper,
		authority:    authority,
		logger:       logger,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetTotalRequested returns the total amount requested by an address
func (k Keeper) GetTotalRequested(ctx context.Context, addr sdk.AccAddress) uint64 {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.GetTotalRequestedKey(addr))
	if bz == nil {
		return 0
	}
	var total math.Uint
	if err := total.Unmarshal(bz); err != nil {
		return 0
	}
	return total.Uint64()
}

// AddToTotalRequested adds an amount to the total requested by an address
func (k Keeper) AddToTotalRequested(ctx context.Context, addr sdk.AccAddress, amount uint64) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	total := k.GetTotalRequested(ctx, addr)
	newTotal := math.NewUint(total + amount)
	bz, err := newTotal.Marshal()
	if err != nil {
		return err
	}
	store.Set(types.GetTotalRequestedKey(addr), bz)
	return nil
}

// AppendRequest stores a request for an address
func (k Keeper) AppendRequest(ctx context.Context, addr sdk.AccAddress, amount uint64, blockHeight int64) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// Create a unique key for this request using address and block height
	key := types.GetRequestKey(addr, blockHeight)

	// Store the amount as big-endian uint64
	amountBytes := sdk.Uint64ToBigEndian(amount)
	store.Set(key, amountBytes)

	return nil
}

// GetRequestsByAddress retrieves all requests made by an address
func (k Keeper) GetRequestsByAddress(ctx context.Context, addr sdk.AccAddress) []types.RequestData {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	var requests []types.RequestData
	prefix := types.GetRequestsKey(addr)

	// Iterate over all keys with the address prefix
	iterator := store.Iterator(prefix, append(prefix, 0xff))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()

		// The key format is: RequestsKeyPrefix (8 bytes) | address (variable) | blockHeight (8 bytes)
		// We need to extract the blockHeight from the last 8 bytes
		if len(key) >= 8 {
			// Get the last 8 bytes for block height
			blockHeightBytes := key[len(key)-8:]
			blockHeight := int64(sdk.BigEndianToUint64(blockHeightBytes))
			amount := sdk.BigEndianToUint64(value)

			requests = append(requests, types.RequestData{
				Amount: amount,
				Height: blockHeight,
			})
		}
	}

	return requests
}
