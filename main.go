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

// make into command line arguments
var (
	in     = filepath.Join(cwd(), "fixtures", "cms.json")
	out    = filepath.Join(cwd(), "src", "images")
	rel    = "/images"
	needle = "https://drive.google.com/file/d/"
)

func Run() error {
	l := Lister{
		needle,
		in,
		make(map[string]Asset),
		struct{}{},
	}

	d := Downloader{
		out,
		rel,
		[]Res{},
	}

	if err := l.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	if err := d.Run(l.Assets); err != nil {
		return fmt.Errorf("%q", err)
	}

	fmt.Printf("%q", l.Assets)

	return nil
}
