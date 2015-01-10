package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var cwd string = filepath.Abs(filepath.Dir(os.Args[0]))

func TestRun(t *testing.T) {

	a := AssetLister{filepath.Join(cwd, "fixtures", "cms.json")}

	if err := a.Run(); err != nil {
		t.Fatalf("%v", err)
	}

	if a.Assets[0] != "helsssl" {
		t.Fatalf("%s", "should = heloo")
	}

}
