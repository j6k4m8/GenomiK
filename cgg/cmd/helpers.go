package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/codegangsta/cli"
)

// Response defines a standard JSON response type for a command.
// Content is generic but should be able to be marshalled by encoding/json.
type Response struct {
	Ok      bool        `json:"ok"`
	Content interface{} `json:"content"`
}

// CLIFunc defines an expected function type for codegangsta/cli.
type CLIFunc func(*cli.Context)

// WrapFunc defines an expected function type for a soon-to-be wrapped function
// for JSON output with codegangsta/cli.
type WrapFunc func(*cli.Context) *Response

// Wrap takes a WrapFunc, calls it with the given context, converts its output
// to JSON, and then outputs that to stdout. This allows these WrapFuncs to be
// used with codegangsta/cli without having to marshal their response
// themselves and while being able to use codegansta/cli contexts without
// modification.
func Wrap(fn WrapFunc) CLIFunc {
	return func(context *cli.Context) {
		resp := fn(context)
		if resp == nil {
			resp = &Response{Ok: false, Content: "Nil response."}
		}
		respJ, err := json.Marshal(resp)
		if err != nil {
			// raw JSON output because marshalling the response struct failed
			// so trying again with an error will probably do the same?
			fmt.Println(`{"ok":false,"content": "Failed to marshall JSON"}`)
			return
		}
		fmt.Println(string(respJ))
	}
}

// ErrorMissingArgument returns a response for a command that is missing an
// argument.
func ErrorMissingArgument() *Response {
	return &Response{
		Ok:      false,
		Content: "Missing argument.",
	}
}

// ErrorOccured wraps a Response object around an error that occurred in a
// command.
func ErrorOccured(e error) *Response {
	return &Response{
		Ok:      false,
		Content: e.Error(),
	}
}
