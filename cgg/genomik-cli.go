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
		{
			Name:    "overlap",
			Aliases: []string{"bbr"},
			Usage: "provide a FASTA file path argument and the overlaps " +
				"will be computed.",
			Action: cmd.Wrap(cmd.Overlap),
		},
	}

	app.Run(os.Args)
}