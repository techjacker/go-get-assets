package main

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"strings"
)

type Asset struct {
	Path string
	Err  error
}
type Lister struct {
	// input
	Needle    string
	InputPath string
	// output
	Data interface{}
	// Assets map[string]struct{}
	Assets map[string]Asset
	// Assets map[string]interface{}
	// Assets    map[string]string
	// Assets    []string
}

func (a *Lister) readFile() ([]byte, error) {
	// read file
	file, err := ioutil.ReadFile(a.InputPath)
	in := bytes.NewReader(file)
	// create buffer of contents
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	return buf.Bytes(), err
}

func (a *Lister) SearchCell(cell string) {
	if strings.Contains(cell, a.Needle) {
		a.Assets[cell] = Asset{}
		// a.Assets[cell] = struct{}{}
		// a.Assets = append(a.Assets, cell)
	}
}

func (a *Lister) Search(cell string) {
	for _, p := range strings.Split(cell, ",") {
		a.SearchCell(strings.TrimSpace(p))
	}
}

func (a *Lister) searchArray(d []interface{}) {
	for _, v := range d {
		switch v.(type) {
		case string:
			a.Search(v.(string))
		case []interface{}:
			a.searchArray(v.([]interface{}))
		case map[string]interface{}:
			a.searchMap(v.(map[string]interface{}))
		}
	}
}

func (a *Lister) searchMap(d map[string]interface{}) {
	for _, v := range d {
		switch v.(type) {
		case string:
			a.Search(v.(string))
		case []interface{}:
			a.searchArray(v.([]interface{}))
		case map[string]interface{}:
			a.searchMap(v.(map[string]interface{}))
		}
	}
}

// search through JSON and create map of assets
func (a *Lister) Run() error {

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
