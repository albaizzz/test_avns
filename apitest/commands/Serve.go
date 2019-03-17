package commands

import (
	"test_avns/apitest/commands/actions"
	"test_avns/apitest/helpers"
	"github.com/urfave/cli"
)

// Serve serves the service
var Serve = cli.Command{
	Name:   "serve",
	Usage:  "Used to run the service",
	Action: actions.RunServer,
	Flags: []cli.Flag{
		helpers.StringFlag("config, c", "configurations/App.yaml", "Custom configuration file path"),
	},
}
