package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultNovaInstanceCost is initially set the same as in wasmd
	DefaultNovaInstanceCost uint64 = 60_000
	// DefaultNovaCompileCost set to a large number for testing
	DefaultNovaCompileCost uint64 = 100
)

// NovaGasRegisterConfig is defaults plus a custom compile amount
func NovaGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultNovaInstanceCost
	gasConfig.CompileCost = DefaultNovaCompileCost

	return gasConfig
}

func NewNovaWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(NovaGasRegisterConfig())
}
