package main

import (
	"os"

	"github.com/Carina-labs/nova/v2/app"
	appconfig "github.com/Carina-labs/nova/v2/app/params"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/tendermint/spm/cosmoscmd"
	tmcmds "github.com/tendermint/tendermint/cmd/tendermint/commands"
)

func main() {
	appconfig.RegisterDenoms()
	cmdOptions := GetWasmCmdOptions()
	cmdOptions = append(cmdOptions, cosmoscmd.AddSubCmd(tmcmds.RollbackStateCmd))
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.NewNovaApp,
		cmdOptions...,
	)

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
