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
	Searcher
	Extracter
}

func NewRenamer(needle string, inputPath string, out string, rel string) *Renamer {

	var r Renamer

	r.Needle = needle
	r.Out = out
	r.Rel = rel

	r.Searcher.InputPath = inputPath
	r.Searcher.SearchCell = r.SearchCell

	r.Extracter.reg = `https://drive.google.com/file/d/(.*)`
	// r.Extracter.reg = `https://drive.google.com/file/d/(\w+.+\w+)`

	return &r
}

func (r *Renamer) SearchCell(cell string) string {
	if strings.Contains(cell, r.Needle) {
		filename := r.Extracter.Gdrive(cell)
		renamedURL := r.Rel + "/" + filename
		return renamedURL
	}
	return cell
}

func (r *Renamer) writeJSONToFile(c interface{}) error {
	contents, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.Out, contents, 0644)
}

func (r *Renamer) Run() error {

	c, err := r.readJSONFromFile()
	if err != nil {
		return err
	}

	r.searchMap(c)

	return r.writeJSONToFile(c)
}
