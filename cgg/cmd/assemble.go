package cmd

import "github.com/codegangsta/cli"

func Assemble(context *cli.Context) *Response {
	path := context.Args().First()
	if path == "" {
		return ErrorMissingArgument()
	}

	unitigs, err := computeUnitigs(path)
	if err != nil {
		return ErrorOccured(err)
	}

	return &Response{
		Ok:      true,
		Content: unitigs,
	}

}
