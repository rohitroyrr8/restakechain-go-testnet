package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxPerAddress = []byte("MaxPerAddress")
	// TODO: Determine the default value
	DefaultMaxPerAddress uint64 = 0
)

var (
	KeyMaxPerRequest = []byte("MaxPerRequest")
	// TODO: Determine the default value
	DefaultMaxPerRequest uint64 = 0
)

var (
	KeyRunning = []byte("Running")
	// TODO: Determine the default value
	DefaultRunning bool = false
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxPerAddress uint64,
	maxPerRequest uint64,
	running bool,
) Params {
	return Params{
		MaxPerAddress: maxPerAddress,
		MaxPerRequest: maxPerRequest,
		Running:       running,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxPerAddress,
		DefaultMaxPerRequest,
		DefaultRunning,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxPerAddress, &p.MaxPerAddress, validateMaxPerAddress),
		paramtypes.NewParamSetPair(KeyMaxPerRequest, &p.MaxPerRequest, validateMaxPerRequest),
		paramtypes.NewParamSetPair(KeyRunning, &p.Running, validateRunning),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxPerAddress(p.MaxPerAddress); err != nil {
		return err
	}

	if err := validateMaxPerRequest(p.MaxPerRequest); err != nil {
		return err
	}

	if err := validateRunning(p.Running); err != nil {
		return err
	}

	return nil
}

// validateMaxPerAddress validates the MaxPerAddress param
func validateMaxPerAddress(v interface{}) error {
	maxPerAddress, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxPerAddress

	return nil
}

// validateMaxPerRequest validates the MaxPerRequest param
func validateMaxPerRequest(v interface{}) error {
	maxPerRequest, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxPerRequest

	return nil
}

// validateRunning validates the Running param
func validateRunning(v interface{}) error {
	running, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = running

	return nil
}
