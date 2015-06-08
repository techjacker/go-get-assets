package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Renamer struct {
	Needle string
	Out    string
	Rel    string
	// Assets map[string]Asset
	Searcher
}

func NewRenamer(needle string, inputPath string, out string, rel string) *Renamer {

	var r Renamer

	r.Needle = needle
	r.Out = out
	r.Rel = rel

	r.Searcher.InputPath = inputPath
	r.Searcher.Data = struct{}{}
	r.Searcher.SearchCell = r.SearchCell

	return &r
}

func (r Renamer) SearchCell(cell string) {
	for _, s := range strings.Split(cell, ",") {
		// r.SearchCell(strings.TrimSpace(s))
		println(strings.TrimSpace(s))
	}
}

func (r Renamer) Run() error {

	var c map[string]interface{}
	contents, err := ioutil.ReadFile(r.InputPath)

	if err != nil {
		return err
	}

	json.Unmarshal(contents, &c)

	err = ioutil.WriteFile(r.Out, contents, 0644)
	return err
}
