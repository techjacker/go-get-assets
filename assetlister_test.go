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
	a.Search("http://themummyofmulberryavenue.com")
	if len(a.Assets) != 0 {
		t.Fatalf("%s", "should not have found any assets")
	}

	a.Assets = []string{}
	assetOne := "http://gdrive.com/2243/"
	a.Search(assetOne)
	if len(a.Assets) != 1 {
		t.Fatalf("%s", "should have found one asset")
	}

	a.Assets = []string{}
	assetArray := "http://gdrive.com/2243/, http://gdrive.com/diffid/"
	a.Search(assetArray)
	if len(a.Assets) != 2 {
		t.Fatalf("%s", "should have found two assets")
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

	// before de duping
	// set in golang??
	if len(a.Assets) != 7 {
		// after dedpuing
		// if len(a.Assets) != 6 {
		t.Fatal("it shd have found 8 assets")
	}

	// t.Logf("%v", a.Data["nophotos"])
	// if a.Assets[0] != "helsssl" {
	// 	t.Fatalf("%s", "should = heloo")
	// }

}
