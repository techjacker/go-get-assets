package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Alister struct {
	Needle    string
	InputPath string
	Data      interface{}
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

func (a *Alister) SearchForAssets(cell string) {

}

// extract a list of assets from the JSON
func (a *Alister) buildAssetList() error {

	d := a.Data.(map[string]interface{})

	for _, v := range d {

		switch v.(type) {
		case string:
			fmt.Println("string:::: " + v.(string))
		case []interface{}:
			fmt.Println("slice:::: ")
		case map[string]interface{}:
			fmt.Println("map:::: ")
		}

	}

	return nil
}

func (a *Alister) Run() error {

	contents, err := a.readFile()

	if err != nil {
		return err
	}

	if err = json.Unmarshal(contents, &a.Data); err != nil {
		return err
	}

	a.buildAssetList()

	// if err = a.buildAssetList(); err != nil {
	// 	return err
	// }
	return err
}
