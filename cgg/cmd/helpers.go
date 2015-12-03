package cmd

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/codegangsta/cli"
)

type Response struct {
	Ok      bool        `json:"ok"`
	Content interface{} `json:"content"`
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

func ErrorMissingArgument() *Response {
	return &Response{
		Ok:      false,
		Content: "Missing argument.",
	}
}

func ErrorOccured(e error) *Response {
	return &Response{
		Ok:      false,
		Content: e.Error(),
	}
}

type StringSet struct {
	contents map[string]int8
	lk       *sync.Mutex
}

func NewStringSet() *StringSet {
	return &StringSet{
		contents: make(map[string]int8),
	}
}

func NewSafeStringSet() *StringSet {
	return &StringSet{
		contents: make(map[string]int8),
		lk:       &sync.Mutex{},
	}
}

func (s *StringSet) IsSafe() bool {
	return s.lk != nil
}

func (s *StringSet) AddContains(value string) bool {
	if s.IsSafe() {
		s.lk.Lock()
		defer s.lk.Unlock()
	}
	if _, exists := s.contents[value]; exists {
		return false
	}
	s.contents[value] = 0
	return true
}

func (s *StringSet) Add(value string) {
	if s.IsSafe() {
		s.lk.Lock()
		defer s.lk.Unlock()
	}
	s.contents[value] = 0
}

func (s *StringSet) Remove(value string) {
	if s.IsSafe() {
		s.lk.Lock()
		defer s.lk.Unlock()
	}
	delete(s.contents, value)
}

func (s *StringSet) Contains(value string) bool {
	s.lk.Lock()
	defer s.lk.Unlock()
	_, exists := s.contents[value]
	return exists
}
