package cmd

import "github.com/codegangsta/cli"

func Hello(c *cli.Context) *Response {
	return &Response{
		Ok:      true,
		Content: "Hello World!",
	}
}
