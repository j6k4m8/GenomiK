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
		{
			Name: "unitig",
			Usage: "provide a FASTA file path argument and the unitigs will " +
				"be computed.",
			Action: cmd.Wrap(cmd.Unitig),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  cmd.OutputFlag,
					Value: "",
					Usage: "optional - provide a path and unitigs will be " +
						"output to it (comma separated).",
				},
				cli.BoolFlag{
					Name:  cmd.PlainTextFlag,
					Usage: "optional - set to output plaintext rather than gzip.",
				},
			},
		},
		{
			Name: "assemble-matrix",
			Usage: "provide a FASTA file path argument and the genome will " +
				"be assembled.",
			Action: cmd.Wrap(cmd.Assemble),
		},
	}

	app.Run(os.Args)
}
