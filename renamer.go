package main

import (
	"io/ioutil"
)

type Renamer struct {
	in     string
	out    string
	rel    string
	needle string
}

func (r Renamer) Run() error {
	contents, err := ioutil.ReadFile(r.in)
	// println(string(contents))
	// println(r.out)
	err = ioutil.WriteFile(r.out, contents, 0644)
	return err
}
