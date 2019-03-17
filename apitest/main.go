package main

import (
	"os"

	"test_avns/apitest/commands"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "test_avns/apitest"
	app.Usage = "Used as service for test_avns/apitest"
	app.UsageText = "[global options] command [arguments]"
	app.Version = "1.0"
	app.Commands = []cli.Command{
		commands.Serve,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
