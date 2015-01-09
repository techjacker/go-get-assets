package main

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"testing"
)

cwd := filepath.Abs(filepath.Dir(os.Args[0]))

func readJSONTest(t *testing.T) {

  a := AssetLister{filepath.Join(cwd, "fixtures", "cms.json")}

  a.LoadInput()

}
