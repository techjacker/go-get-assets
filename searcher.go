package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Searcher struct {
	InputPath  string
	SearchCell func(string) string
}

func (s *Searcher) readJSONFromFile() (map[string]interface{}, error) {

	contents, err := ioutil.ReadFile(s.InputPath)
	if err != nil {
		return nil, err
	}

	var c map[string]interface{}
	json.Unmarshal(contents, &c)

	return c, err
}

func (s *Searcher) Search(cell string) string {
	var res string
	for _, p := range strings.Split(cell, ",") {
		res += s.SearchCell(strings.TrimSpace(p)) + ","
	}
	return strings.TrimRight(res, ",")
}

func (s *Searcher) searchArray(d []interface{}) {
	for _, v := range d {
		switch v.(type) {
		case string:
			s.Search(v.(string))
		case []interface{}:
			s.searchArray(v.([]interface{}))
		case map[string]interface{}:
			s.searchMap(v.(map[string]interface{}))
		}
	}
}

func (s *Searcher) searchMap(d map[string]interface{}) map[string]interface{} {
	for k, v := range d {
		switch v.(type) {
		case string:
			newStr := s.Search(v.(string))
			d[k] = newStr
		case []interface{}:
			s.searchArray(v.([]interface{}))
		case map[string]interface{}:
			s.searchMap(v.(map[string]interface{}))
		}
	}
	return d
}
