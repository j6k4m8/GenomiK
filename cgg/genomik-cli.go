package main

import (
	"fmt"
	"os"

	"github.com/j6k4m8/cg/cgg/cmd"

	"github.com/codegangsta/cli"
)

func main() {
	fmt.Println(os.Args)
	app := cli.NewApp()
	app.Name = "genomik-cli"

	gzFlag := cli.BoolFlag{
		Name: cmd.GZipFlag,
		Usage: "set this flag to indicated that the input FASTA file is " +
			"gzipped.",
	}

	app.Commands = []cli.Command{
		{
			Name:    "hello",
			Aliases: []string{"hi"},
			Usage:   "say hi for sanity",
			Action:  cmd.Wrap(cmd.Hello),
			Flags:   []cli.Flag{gzFlag},
		},
		{
			Name:    "overlap",
			Aliases: []string{"bbr"},
			Usage: "provide a FASTA file path argument and the overlaps " +
				"will be computed.",
			Action: cmd.Wrap(cmd.Overlap),
			Flags:  []cli.Flag{gzFlag},
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
				gzFlag,
			},
		},
		{
			Name: "assemble-matrix",
			Usage: "provide a FASTA file path argument and the genome will " +
				"be assembled.",
			Action: cmd.Wrap(cmd.Assemble),
			Flags:  []cli.Flag{gzFlag},
		},
	}

	app.Run(os.Args)
}
