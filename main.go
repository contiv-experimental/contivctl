package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/contiv/netplugin/netctl"
	"github.com/contiv/volplugin/volcli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Central command for both netplugin and volplugin"
	app.Version = ""
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name:        "net",
			Usage:       "Control netplugin",
			Flags:       netctl.NetmasterFlags,
			Subcommands: netctl.Commands,
		},
		{
			Name:        "vol",
			Usage:       "Control volplugin",
			Flags:       volcli.VolmasterFlags,
			Subcommands: volcli.Commands,
		},
	}

	app.Run(os.Args)
}
