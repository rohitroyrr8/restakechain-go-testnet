package types

// RequestData represents a single faucet request with amount and block height
type RequestData struct {
	Amount uint64
	Height int64
}
