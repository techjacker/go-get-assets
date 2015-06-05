package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Renamer struct {
	in     string
	out    string
	rel    string
	needle string
}

func (r Renamer) Run() error {

	var c map[string]interface{}
	contents, err := ioutil.ReadFile(r.in)

	if err != nil {
		return err
	}

	json.Unmarshal(contents, &c)

	// for _, v := range c {
	// 	switch v.(type) {
	// 	case string:
	// 		println(v.(string))
	// 	}
	// }

	// println(string(contents))
	// println(r.out)
	err = ioutil.WriteFile(r.out, contents, 0644)
	return err
}

func (r Renamer) Search(cell string) {
	for _, s := range strings.Split(cell, ",") {
		// r.SearchCell(strings.TrimSpace(s))
		println(strings.TrimSpace(s))
	}
}
