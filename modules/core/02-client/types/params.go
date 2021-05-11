package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/ibc-go/modules/core/exported"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// DefaultAllowedClients are "06-solomachine" and "07-tendermint"
	DefaultAllowedClients = []string{exported.Solomachine, exported.Tendermint, "wasm_dummy"}

	DefaultWASMClientEnabled = false

	// KeyAllowedClients is store's key for AllowedClients Params
	KeyAllowedClients = []byte("AllowedClients")

	KeyWasmClientsEnabled = []byte("WasmClientsEnabled")
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the ibc transfer module
func NewParams(wasmClientAllowed bool, allowedClients ...string) Params {
	return Params{
		WasmClientsEnabled: wasmClientAllowed,
		AllowedClients:     allowedClients,
	}
}

// DefaultParams is the default parameter configuration for the ibc-transfer module
func DefaultParams() Params {
	return NewParams(DefaultWASMClientEnabled, DefaultAllowedClients...)
}

// Validate all ibc-transfer module parameters
func (p Params) Validate() error {
	err := validateClients(p.AllowedClients)
	if err != nil {
		return err
	}

	return validateWasmClientEnabledFlag(p.WasmClientsEnabled)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAllowedClients, p.AllowedClients, validateClients),
		paramtypes.NewParamSetPair(KeyWasmClientsEnabled, p.WasmClientsEnabled, validateWasmClientEnabledFlag),
	}
}

// IsAllowedClient checks if the given client type is registered on the allowlist.
func (p Params) IsAllowedClient(clientType string) bool {
	for _, allowedClient := range p.AllowedClients {
		if allowedClient == clientType {
			return true
		}
	}
	return false
}

func validateWasmClientEnabledFlag(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateClients(i interface{}) error {
	clients, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for i, clientType := range clients {
		if strings.TrimSpace(clientType) == "" {
			return fmt.Errorf("client type %d cannot be blank", i)
		}
	}

	return nil
}
