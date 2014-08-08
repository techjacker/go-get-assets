package main

import (
	"fmt"
	"os"
)

var assetLister = AssetLister{}

// type AssetGetter struct {
// 	AssetLister
// 	AssetDownloader
// }

// $ go-get-assets path-to-data.json outPutDirPath
func main() {
	jsonPath := os.Args[1]
	fmt.Println(jsonPath)
	// var assetLister = AssetDownloader()
}
