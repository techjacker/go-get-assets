package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Alister struct {
	InputPath string
	Data      map[string]interface{}
	Assets    []string
}

func (a *Alister) readFile() ([]byte, error) {
	// read file
	file, err := ioutil.ReadFile(a.InputPath)
	in := bytes.NewReader(file)
	// create buffer of contents
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	return buf.Bytes(), err
}

// extract a list of assets from the JSON
func (a *Alister) buildAssetList() error {
	// a.Assets{"hello", "world"}
	return nil
}

func (a *Alister) Run() error {

	contents, err := a.readFile()

	if err != nil {
		return err
	}

	err = json.Unmarshal(contents, &a.Data)

	// if err = a.marshallInput(bytes.NewReader(file)); err != nil {
	// 	return err
	// }

	// if err = a.buildAssetList(); err != nil {
	// 	return err
	// }
	return err
}
