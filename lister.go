package main

import (
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

func (l *Lister) Run() error {
	c, err := l.readJSONFromFile()
	if err != nil {
		return err
	}
	// extract a list of assets from the JSON
	l.searchMap(c)
	return err
}
