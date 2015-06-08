package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRewriteUrlsInJson(t *testing.T) {

	of, err := ioutil.TempFile("", "renamer-test")
	defer of.Close()
	if err != nil {
		t.Fatal(err)
	}

	var (
		in        = filepath.Join(Cwd(), "fixtures", "cms.renamer.input.json")
		expected  = filepath.Join(Cwd(), "fixtures", "cms.renamer.expected.json")
		out       = of.Name()
		rel       = "/images"
		needle    = "https://drive.google.com/file/d/"
		want, got map[string]interface{}
	)

	r := NewRenamer(
		needle,
		in,
		out,
		rel,
	)

	if err = r.Run(); err != nil {
		t.Fatal(err)
	}

	w, err := ioutil.ReadFile(expected)
	o, err := ioutil.ReadFile(r.Out)
	json.Unmarshal(w, &want)
	json.Unmarshal(o, &got)
	// t.Logf("%v", o["mapofphotos"].(map[string]interface{})["photourl"])
	if !reflect.DeepEqual(want, got) {
		t.Fatal("want does not equal got")
	}
}
