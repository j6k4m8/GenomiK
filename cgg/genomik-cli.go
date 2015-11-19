package main

import (
	"os"

	"cg/cgg/handlers"

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
			Action:  handlers.Wrap(handlers.Hello),
		},
	}

	app.Run(os.Args)
}
