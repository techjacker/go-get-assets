package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

type assetLister interface {
	readJson() [string]string
	getListOfAssets() []string
}

type assetReader interface {
	readJson() [string]string
	getListOfAssets() []string
}
