package main

import (
	"os"

	"github.com/j6k4m8/cg/cgg/cmd"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "genomik-cli"

	app.Commands = []cli.Command{
		{
			Name:    "hello",
			Aliases: []string{"hi"},
			Usage:   "say hi for sanity",
			Action:  cmd.Wrap(cmd.Hello),
		},
	}

	app.Run(os.Args)
}
