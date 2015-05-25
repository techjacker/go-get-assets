package main

import (
	"github.com/clbanning/mxj"
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

	gotMap, err := mxj.NewMapJson(got)
	wantMap, err := mxj.NewMapJson(want)

	if !reflect.DeepEqual(gotMap, wantMap) {
		t.Fatal("want does not equal got")
	}

	// v, err := m.ValuesForKey("nophotos")
	// println(v[0].(string))
}
