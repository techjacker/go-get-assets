package main

import (
	"fmt"
	// "os"
)

// a.InputPath = filepath.Join(cwd, "fixtures", "cms.json")

func main() {
	g, err := gdrive.New()
	a := Lister{
		"/home/andy/lib/modules/go/src/github.com/techjacker/go-get-assets/fixtures/cms.json",
		"http://gdrive.com",
		make(map[string]struct{}),
	}

	if err := a.Run(); err != nil {
		fmt.Printf("%v", err)
	}

}
