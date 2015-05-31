package main

import (
	"encoding/json"
	"strings"
)

type Asset struct {
	Path string
	Err  error
}

type Lister struct {
	Assets map[string]Asset
	JsonLooper
	// SearcherARunner
}

func (a *Lister) SearchCell(cell string) {
	if strings.Contains(cell, a.Needle) {
		a.Assets[cell] = Asset{}
		// a.Assets[cell] = struct{}{}
		// a.Assets = append(a.Assets, cell)
	}
}

// search through JSON and create map of assets
func (a *Lister) Run() error {

	// data := struct{}{}
	// data := interface{}{}
	contents, err := a.readFile()

	if err != nil {
		return err
	}

	if err = json.Unmarshal(contents, &a.Data); err != nil {
		return err
	}

	// extract a list of assets from the JSON
	a.searchMap(a.Data.(map[string]interface{}))

	return nil
}
