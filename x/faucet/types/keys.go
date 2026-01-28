package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "faucet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_faucet"
)

var (
	ParamsKey               = []byte("p_faucet")
	TotalRequestedKeyPrefix = []byte("t_requested")
	RequestsKeyPrefix       = []byte("requests")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetTotalRequestedKey returns the store key for total requested amount by an address
func GetTotalRequestedKey(addr sdk.AccAddress) []byte {
	return append(TotalRequestedKeyPrefix, addr.Bytes()...)
}

// GetRequestsKey returns the store key prefix for all requests by an address
func GetRequestsKey(addr sdk.AccAddress) []byte {
	return append(RequestsKeyPrefix, addr.Bytes()...)
}

// GetRequestKey returns the store key for a specific request by address and block height
func GetRequestKey(addr sdk.AccAddress, blockHeight int64) []byte {
	key := append(GetRequestsKey(addr), sdk.Uint64ToBigEndian(uint64(blockHeight))...)
	return key
}
