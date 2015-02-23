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
	var a Lister

	a.Needle = "http://gdrive.com"
	a.Search("http://themummyofmulberryavenue.com")
	if len(a.Assets) != 0 {
		t.Fatalf("%s", "should not have found any assets")
	}

	a.Assets = map[string]Asset{}
	assetOne := "http://gdrive.com/2243/"
	a.Search(assetOne)
	if len(a.Assets) != 1 {
		t.Fatalf("%s", "should have found one asset")
	}

	a.Assets = map[string]Asset{}
	// a.Assets = []string{}
	assetArray := "http://gdrive.com/2243/, http://gdrive.com/diffid/"
	a.Search(assetArray)
	if len(a.Assets) != 2 {
		t.Fatalf("%s", "should have found two assets")
	}

}

func TestRun(t *testing.T) {
	// t.Errorf("%v", &cwd)

	var a Lister
	// a.InputPath = filepath.Join(cwd, "fixtures", "cms.json")
	a.InputPath = "/home/andy/go/src/github.com/techjacker/go-get-assets/fixtures/cms.json"
	a.Needle = "http://gdrive.com"
	a.Assets = make(map[string]Asset)

	if err := a.Run(); err != nil {
		t.Fatalf("%v", err)
	}

	if len(a.Assets) != 6 {
		t.Fatalf("%s%d", "# assets = ", len(a.Assets))
		t.Fatalf("\n\n%v\n\n", a.Assets)
	}
	// t.Logf("%q", a.Assets)
	// map["http://gdrive.com/traffic.jpg":{} "http://gdrive.com/diff.jpg":{} "http://gdrive.com/city.jpg":{} "http://gdrive.com/sf.jpg":{} "http://gdrive.com/highway.jpg":{} "http://gdrive.com/freedom.jpg":{}]
}
