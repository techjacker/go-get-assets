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

	var err error

	l := NewLister(needle, in)
	if err = l.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	d := NewDownloader(imagesDir, rel)
	if err = d.Run(l.Assets); err != nil {
		return fmt.Errorf("%q", err)
	}
	fmt.Printf("%q", l.Assets)

	r := NewRenamer(needle, in, out, rel)
	if err = r.Run(); err != nil {
		return fmt.Errorf("%q", err)
	}

	return err
}
