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
	Needle string
	Assets map[string]Asset
	Searcher
}

func NewLister(needle string, inputPath string) *Lister {
	var l Lister
	l.Needle = needle
	l.Assets = map[string]Asset{}
	l.Searcher.InputPath = inputPath
	l.Searcher.Data = struct{}{}
	l.Searcher.SearchCell = l.SearchCell
	return &l
}

func (l *Lister) SearchCell(cell string) string {
	if strings.Contains(cell, l.Needle) {
		l.Assets[cell] = Asset{}
		// l.Assets[cell] = struct{}{}
		// l.Assets = append(l.Assets, cell)
	}
	return cell
}

// search through JSON and create map of assets
func (l *Lister) Run() error {

	// data := struct{}{}
	// data := interface{}{}
	contents, err := l.readFile()

	if err != nil {
		return err
	}

	if err = json.Unmarshal(contents, &l.Data); err != nil {
		return err
	}

	// extract a list of assets from the JSON
	l.searchMap(l.Data.(map[string]interface{}))

	return nil
}
