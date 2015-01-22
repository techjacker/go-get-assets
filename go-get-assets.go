package main

import (
	"fmt"
	// "os"
)

func main() {
	var a Lister
	// a.InputPath = filepath.Join(cwd, "fixtures", "cms.json")
	a.InputPath = "/home/andy/lib/modules/go/src/github.com/techjacker/go-get-assets/fixtures/cms.json"
	a.Needle = "http://gdrive.com"
	a.Assets = make(map[string]struct{})

	if err := a.Run(); err != nil {
		fmt.Printf("%v", err)
	}
}
