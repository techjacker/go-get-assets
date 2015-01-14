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

func TestSearchForAssets(t *testing.T) {
	var a Alister
	a.Needle = "http://gdrive.com"
	a.SearchForAssets("http://themummyofmulberryavenue.com")
	if len(a.Assets) != 0 {
		t.Fatalf("%s", "should not have found any assets")
	}
}

func TestRun(t *testing.T) {
	// t.Errorf("%v", &cwd)

	var a Alister
	// a.InputPath = filepath.Join(cwd, "fixtures", "cms.json")
	a.InputPath = "/home/andy/lib/modules/go/src/github.com/techjacker/go-get-assets/fixtures/cms.json"
	a.Needle = "http://gdrive.com"

	if err := a.Run(); err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("%s%d", "# assets = ", len(a.Assets))

	// t.Logf("%v", a.Data["nophotos"])
	// if a.Assets[0] != "helsssl" {
	// 	t.Fatalf("%s", "should = heloo")
	// }

}
