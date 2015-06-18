package main

import (
	"path/filepath"
)

var gDriveURL string = "https://drive.google.com/file/d"

// $ go-get-assets <input-json> <download-dir> <rewrite-urls-stem> <output-json>

// make into command line arguments
var (
	in        = filepath.Join(Cwd(), "fixtures", "cms.json")
	out       = filepath.Join(Cwd(), "fixtures", "cms.downloaded.json")
	imagesDir = filepath.Join(Cwd(), "src", "images")
	rel       = "/images"
	needle    = gDriveURL
)

func New(needle, imagesDir, rel, out string) *GoGetAsseter {
	return &GoGetAsseter{
		l: NewLister(needle, in),
		d: NewDownloader(imagesDir, rel),
		r: NewRenamer(needle, in, out, rel),
	}
}
