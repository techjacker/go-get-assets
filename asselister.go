package main

import (
	"bytes"
	"io/ioutil"
)

// "fmt"

type AssetLister struct {
	InputPath string
	Data      map[string]interface{}
	AssetList []string
}

// func (a *AssetLister) readJSON(pathToJson string, map[string]interface{})  {
func (a *AssetLister) LoadInput() {
	file, err := ioutil.ReadFile(InputPath)
	if err != nil {
		return err
	}
	a.marshallInput(bytes.NewReader(file), a.Data)
}

func (a AssetLister) marshallInput(in io.Reader, d map[string]interface) {

}

// // extract a list of assets from the JSON
// func (a *AssetLister) BuildAssetList(interface{}) []string {
// 	return nil
// 	// fmt.println
// }
