package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRewriteUrlsInJson(t *testing.T) {

	type CmsExample struct {
		Nophotos string `json:"nophotos"`
		// Mapofphotos struct {
		// 	Photourl   string
		// 	Nastyarray []string
		// 	Notaphoto  string
		// }
		// Clients []struct {
		// 	Name     string
		// 	Photourl string
		// }
		// Projects []struct {
		// 	Descriptionsub     string
		// 	Descriptionsublink string
		// 	// Photourl           []interface{}
		// 	Photourl string
		// }
	}

	outFile, err := ioutil.TempFile("", "renamer-test")
	if err != nil {
		t.Fatal(err)
	}

	var (
		out    = outFile.Name()
		in     = filepath.Join(Cwd(), "fixtures", "cms.strong.json")
		rel    = "/images"
		needle = "https://drive.google.com/file/d/"
		// result CmsExample
	)

	r := Renamer{
		in,
		out,
		rel,
		needle,
	}

	err = r.Run()
	if err != nil {
		t.Fatal(err)
	}

	result := &CmsExample{}
	// outFileModified, err := os.OpenFile(outFile.Name(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	outFileModified, err := os.OpenFile(outFile.Name(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		t.Fatal(err)
	}
	// err = json.NewDecoder(outFile).Decode(result)
	err = json.NewDecoder(outFileModified).Decode(result)
	// err = mapstructure.Decode(outFile, result)
	defer outFile.Close()
	defer outFileModified.Close()
	// mapstructure.Decode(outFileModified, result)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
	t.Logf("%v", result)
	// results := []Res{
	// 	Res{
	// 		[]byte{},
	// 		nil,
	// 		"http://gdrive.com/diff.jpg",
	// 		"diff",
	// 	},
	// }

	// err = r.Run(results)

}
