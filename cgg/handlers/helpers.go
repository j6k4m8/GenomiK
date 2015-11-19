package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/codegangsta/cli"
)

type Response struct {
	Ok      bool   `json:"ok"`
	Content string `json:"content"`
}

type CLIFunc func(*cli.Context)
type WrapFunc func(*cli.Context) *Response

// Wrap takes a WrapFunc, calls it with the given context, converts its output
// to JSON, and then outputs that to stdout. This allows these WrapFuncs to be
// used with codegangsta/cli without having to marshal their response
// themselves and while being able to use codegansta/cli contexts without
// modification.
func Wrap(fn WrapFunc) CLIFunc {
	return func(context *cli.Context) {
		resp := fn(context)
		respJ, err := json.Marshal(resp)
		if err != nil {
			// raw JSON output because marshalling the response struct failed
			// so trying again with an error will probably do the same?
			fmt.Println(`{"ok":false,"content": "Failed to marshall JSON"}`)
		}
		fmt.Println(string(respJ))
	}
}
