package main

import (
	"encoding/json"
	"fmt"
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

func (r Renamer) readJSONFromFile(inputPath string) (map[string]interface{}, error) {

	contents, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	var c map[string]interface{}
	json.Unmarshal(contents, &c)

	return c, err
}

func (r Renamer) Run() error {

	contents, err := ioutil.ReadFile(r.InputPath)
	if err != nil {
		return err
	}

	c, err := r.readJSONFromFile(r.InputPath)

	println(c)
	fmt.Printf("%v", c)
	r.searchMap(c)

	err = ioutil.WriteFile(r.Out, contents, 0644)
	return err
}
