package main

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
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

func ExtractGdriveFilename(url string) string {
	// idReg := regexp.MustCompile(`https://drive.google.com/file/d/(\w+.+\w+)`)
	idReg := regexp.MustCompile(`https://drive.google.com/file/d/(.*)`)
	id := idReg.FindStringSubmatch(url)
	// didn't find a match
	if len(id) < 2 {
		return ""
	}
	return id[1]
}

func (r Renamer) SearchCell(cell string) string {
	if strings.Contains(cell, r.Needle) {
		renamedURL := r.Rel + "/" + ExtractGdriveFilename(cell)
		return renamedURL
	}
	return cell
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

func (r Renamer) writeJSONToFile(c interface{}) error {
	contents, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.Out, contents, 0644)
}

func (r Renamer) Run() error {
	c, err := r.readJSONFromFile(r.InputPath)
	if err != nil {
		return err
	}
	r.searchMap(c)
	// fmt.Printf("%v", c)
	return r.writeJSONToFile(c)
}
