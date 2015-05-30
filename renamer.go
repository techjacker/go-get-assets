package main

import (
	"encoding/json"
	"io/ioutil"
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
	json.Unmarshal(contents, &c)

	for _, v := range c {
		switch v.(type) {
		case string:
			println(v.(string))
		}
	}

	// println(string(contents))
	// println(r.out)
	err = ioutil.WriteFile(r.out, contents, 0644)
	return err
}
