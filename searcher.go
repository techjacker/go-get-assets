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

type JsonLooper struct {
	Needle     string
	InputPath  string
	Data       interface{}
	SearchCell func(string)
}

func (a *JsonLooper) readFile() ([]byte, error) {
	// read file
	file, err := ioutil.ReadFile(a.InputPath)
	in := bytes.NewReader(file)
	// create buffer of contents
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	return buf.Bytes(), err
}

func (a *JsonLooper) Search(cell string) {
	for _, p := range strings.Split(cell, ",") {
		a.SearchCell(strings.TrimSpace(p))
	}
}

func (a *JsonLooper) searchArray(d []interface{}) {
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

func (a *JsonLooper) searchMap(d map[string]interface{}) {
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
