package main

import (
	"fmt"
)

type AssetLister interface {
	// maybe there is a filepath type???
	// import ("filepathpackage") filepathpackage.path type????
	// return array of download link strings
	readJson(pathToJson string) []string
	getListOfAssets() []string
}
