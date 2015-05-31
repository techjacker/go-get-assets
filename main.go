package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Cwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

// make into command line arguments
var (
	in        = filepath.Join(Cwd(), "fixtures", "cms.json")
	out       = filepath.Join(Cwd(), "fixtures", "cms.downloaded.json")
	imagesDir = filepath.Join(Cwd(), "src", "images")
	rel       = "/images"
	needle    = "https://drive.google.com/file/d/"
)

func Run() error {
	var l Lister

	l.Assets = make(map[string]Asset)
	l.Needle = needle
	l.InputPath = in
	l.Data = struct{}{}

	// l := Lister{
	// 	make(map[string]Asset),
	// 	JsonLooper{needle, in, struct{}{}, Lister.SearchCell},
	// 	// needle, in, struct{}{},
	// }

	d := Downloader{
		imagesDir,
		rel,
		[]Res{},
	}

	// r := Renamer{
	// 	in,
	// 	out,
	// 	rel,
	// 	needle,
	// }

	if err := l.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	if err := d.Run(l.Assets); err != nil {
		return fmt.Errorf("%q", err)
	}

	// if err := r.Run(d.Results, io.Writer{}); err != nil {
	// 	return fmt.Errorf("%q", err)
	// }

	fmt.Printf("%q", l.Assets)

	return nil
}
