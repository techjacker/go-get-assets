package main

import (
	"path/filepath"
	// "encoding/json"
	"testing"
)

func TestRewriteUrlsInJson(t *testing.T) {
	var (
		in     = filepath.Join(Cwd(), "fixtures", "cms.json")
		out    = filepath.Join(Cwd(), "fixtures", "cms.downloaded.json")
		rel    = "/images"
		needle = "https://drive.google.com/file/d/"
	)

	r := Renamer{
		in,
		out,
		rel,
		needle,
	}

	results := []Res{
		Res{
			[]byte{},
			nil,
			"http://gdrive.com/diff.jpg",
			"diff",
		},
	}

	err := r.Run(results)

	if err != nil {
		t.Fatal(err)
	}

	// if v["mapofphotos"]["photourl"] != "/images/diff" {
	// 	t.Fatal("got:", v["mapofphotos"]["photourl"])
	// 	t.Fatal("want:", "/images/diff")
	// }
}
