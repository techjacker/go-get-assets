package main

import (
	"strings"
	// "encoding/json"
	"io/ioutil"
	// "os"
	"path/filepath"
	"testing"
)

func TestRewriteUrlsInJson(t *testing.T) {

	of, err := ioutil.TempFile("", "renamer-test")
	defer of.Close()
	if err != nil {
		t.Fatal(err)
	}

	var (
		in     = filepath.Join(Cwd(), "fixtures", "cms.strong.json")
		out    = of.Name()
		rel    = "/images"
		needle = "https://drive.google.com/file/d/"
	)

	r := Renamer{
		in,
		out,
		rel,
		needle,
	}

	if err = r.Run(); err != nil {
		t.Fatal(err)
	}

	got, err := ioutil.ReadFile(r.out)
	want, err := ioutil.ReadFile(r.in)
	println(string(got))
	println(string(want))

	if strings.TrimSpace(string(got)) != strings.TrimSpace(string(want)) {
		t.Fatal("in does not equal out")
	}

}
