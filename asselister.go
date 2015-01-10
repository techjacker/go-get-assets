package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// "fmt"

type Alister struct {
	InputPath string
	Data      map[string]interface{}
	Assets    []string
}

func (a Alister) loadInput() (io.Reader, error) {
	return ioutil.ReadFile(InputPath)
}

func (a *Alister) marshallInput(in io.Reader) error {
	buf := bytes.NewReader(bytes.Buffer)
	buf.ReadFrom(in)
	return json.Unmarshal(in, &a.Data)
}

// extract a list of assets from the JSON
func (a *Alister) buildAssetList() error {
	a.Assets{"hello", "world"}
	return nil
}

func (a Alister) Run() error {

	if file, err := a.loadInput(); err != nil {
		return err
	}

	if err = a.marshallInput(bytes.NewReader(file)); err != nil {
		return err
	}

	if err = a.buildAssetList(); err != nil {
		return err
	}
}
