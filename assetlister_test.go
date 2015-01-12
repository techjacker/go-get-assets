package main

import (
	// "flag"
	// "path/filepath"
	"testing"
)

// func init() {
// 	flag.Parse()
// }

// var (
// 	cwd = flag.String("cwd", "", "set cwd")
// )

func TestRun(t *testing.T) {
	// t.Errorf("%v", &cwd)

	var a Alister
	// a.InputPath = filepath.Join(cwd, "fixtures", "cms.json")
	a.InputPath = "/home/andy/lib/modules/go/src/github.com/techjacker/go-get-assets/fixtures/cms.json"

	if err := a.Run(); err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("%v", a.Data["nophotos"])
	// if a.Assets[0] != "helsssl" {
	// 	t.Fatalf("%s", "should = heloo")
	// }

}
