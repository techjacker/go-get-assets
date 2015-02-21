package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func cwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

// in     = "/home/andy/lib/modules/go/src/github.com/techjacker/go-get-assets/fixtures/cms.json"
// make into command line arguments
var (
	in     = filepath.Join(cwd(), "fixtures", "cms.json")
	out    = filepath.Join(cwd(), "src", "images")
	rel    = "/images"
	needle = "https://drive.google.com/file/d/"
)

func Run() error {
	var l Lister

	l.Needle = needle
	l.InputPath = in

	if err := l.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("%q", l.Assets)

	d := Downloader{
		out,
		rel,
		l.Assets,
	}

	if err := d.Run(); err != nil {
		return fmt.Errorf("%q", err)
	}

	return nil
}
