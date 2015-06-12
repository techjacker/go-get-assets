package main

import (
	"bytes"

	// "fmt"
	"io/ioutil"
	"strings"
)

type SearcherRunner interface {
	SearchCell(string)
	Run() error
}

type Searcher struct {
	InputPath  string
	Data       interface{}
	SearchCell func(string) string
}

func (s *Searcher) readFile() ([]byte, error) {
	// read file
	file, err := ioutil.ReadFile(s.InputPath)
	in := bytes.NewReader(file)
	// create buffer of contents
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	return buf.Bytes(), err
}

func (s *Searcher) Search(cell string) string {
	var res string
	for _, p := range strings.Split(cell, ",") {
		res += s.SearchCell(strings.TrimSpace(p))
	}
	return res
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
			// println(newStr)
			d[k] = newStr
			// println(d[k])
		case []interface{}:
			s.searchArray(v.([]interface{}))
		case map[string]interface{}:
			s.searchMap(v.(map[string]interface{}))
		}
	}
	return d
}
