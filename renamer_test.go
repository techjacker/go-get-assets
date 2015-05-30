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
		in        = filepath.Join(Cwd(), "fixtures", "cms.strong.json")
		out       = of.Name()
		rel       = "/images"
		needle    = "https://drive.google.com/file/d/"
		want, got map[string]interface{}
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

	i, err := ioutil.ReadFile(r.in)
	o, err := ioutil.ReadFile(r.out)
	json.Unmarshal(i, &want)
	json.Unmarshal(o, &got)
	// t.Logf("%v", o["mapofphotos"].(map[string]interface{})["photourl"])
	if !reflect.DeepEqual(want, got) {
		t.Fatal("want does not equal got")
	}
}
