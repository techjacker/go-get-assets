package main

import (
	"path/filepath"
	"testing"
)

func TestSearchForAssets(t *testing.T) {

	l := NewLister("http://gdrive.com", "")

	l.Search("http://themummyofmulberryavenue.com")
	if len(l.Assets) != 0 {
		t.Fatalf("%s", "should not have found any assets")
	}

	assetOne := "http://gdrive.com/2243/"
	l.Search(assetOne)
	if len(l.Assets) != 1 {
		t.Fatalf("%s", "should have found one asset")
	}

	l.Assets = map[string]Asset{}
	assetArray := "http://gdrive.com/2243/, http://gdrive.com/diffid/"
	l.Search(assetArray)
	if len(l.Assets) != 2 {
		t.Fatalf("%s", "should have found two assets")
	}

}

func TestRun(t *testing.T) {

	var (
		inputPath string  = filepath.Join(Cwd(), "fixtures", "cms.json")
		l         *Lister = NewLister("http://gdrive.com", inputPath)
	)

	if err := l.Run(); err != nil {
		t.Fatalf("%v", err)
	}

	if len(l.Assets) != 6 {
		t.Fatalf("%s%d", "# assets = ", len(l.Assets))
		t.Fatalf("\n\n%v\n\n", l.Assets)
	}
	// t.Logf("%q", a.Assets)
	// map["http://gdrive.com/traffic.jpg":{} "http://gdrive.com/diff.jpg":{} "http://gdrive.com/city.jpg":{} "http://gdrive.com/sf.jpg":{} "http://gdrive.com/highway.jpg":{} "http://gdrive.com/freedom.jpg":{}]
}
