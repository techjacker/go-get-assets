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

	a.Assets = map[string]struct{}{}
	assetOne := "http://gdrive.com/2243/"
	a.Search(assetOne)
	if len(a.Assets) != 1 {
		t.Fatalf("%s", "should have found one asset")
	}

	a.Assets = map[string]struct{}{}
	// a.Assets = []string{}
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
	a.Assets = make(map[string]struct{})

	if err := a.Run(); err != nil {
		t.Fatalf("%v", err)
	}

	if len(a.Assets) != 6 {
		t.Fatalf("%s%d", "# assets = ", len(a.Assets))
		t.Fatalf("\n\n%v\n\n", a.Assets)
	}
}
