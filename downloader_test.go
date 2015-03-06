package main

import (
	"strings"
	// "errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	// "os"
	// "strings"
	"testing"
)

var (
	tId      = "0ByPfUp1fLihSNm5SSjZoalhPQ3M"
	tIn      = "https://drive.google.com/file/d/" + tId + "/view?usp=sharing"
	tRel     = "/"
	eDir     = "src/images/"
	eRelPath = tRel + tId
	ePath    = eDir + tId
	eUrl     = "https://googledrive.com/host/" + tId
)

// func TestExtractIdMissing(t *testing.T) {
// bad string goes here

func TestExtractId(t *testing.T) {
	var d Downloader
	if d.ExtractId(tIn) != tId {
		t.Fatal("got:", d.ExtractId(tIn))
	}
}

func TestCreateDestPath(t *testing.T) {
	var d Downloader
	d.OutputDir = eDir
	got := d.CreateDestPath(tId)
	if got != ePath {
		t.Fatal("got:", got)
	}
}

func TestCreateRelPath(t *testing.T) {
	var d Downloader
	d.RelativePath = tRel
	if d.CreateRelPath(tId) != eRelPath {
		t.Fatal("got:", d.CreateRelPath(tId))
	}
}

func TestCreateTargetUrl(t *testing.T) {
	var d Downloader
	if d.CreateTargetUrl(tId) != eUrl {
		t.Fatal("got:", d.CreateTargetUrl(tId))
	}
}

func TestRewriteUrlsInJson(t *testing.T) {
	var (
		d Downloader
		// v interface{}
	)

	d.RelativePath = "/images"
	d.Results = []Res{
		Res{
			[]byte{},
			nil,
			"http://gdrive.com/diff.jpg",
			"diff",
		},
	}

	// input := []byte(`{
	input := map[string]interface{}{
		"mapofphotos": map[string]interface{}{
			"photourl":   "http://gdrive.com/diff.jpg",
			"nastyarray": "sdfsd",
			"noaphoto":   "sdfsd",
		},
	}
	// }`)

	// json.Unmarshal(input, &v)
	// d.RewriteUrlsInJson(&v)

	t.Log(input["mapofphotos"])
	t.Log(input["mapofphotos"].(map[string]interface{})["photourl"])

	// if v["mapofphotos"]["photourl"] != "/images/diff" {
	// 	t.Fatal("got:", v["mapofphotos"]["photourl"])
	// 	t.Fatal("want:", "/images/diff")
	// }
}

func TestDownload(t *testing.T) {
	var (
		d        Downloader
		res      Res
		body     = "hello"
		id       = "12345"
		origUrl  = "http://sdffdsfsd.com"
		chanDown = make(chan Res, 5)
		done     = make(chan bool)
	)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	go func() {
		res = <-chanDown // block, waiting for res value to be populated
		done <- true     // we're done
	}()
	go d.Download(ts.URL, origUrl, id, chanDown)
	<-done // wait for output to be received by res

	if res.Err != nil {
		t.Error("error is not nil")
		t.Error(res.Err)
	}
	if strings.TrimSpace(string(res.Data)) != body {
		t.Error("data wrong")
		t.Error(res.Data)
	}
	if res.Url != origUrl {
		t.Error("origUrl not passed")
		t.Error(origUrl)
	}
	if res.Id != id {
		t.Error("id not passed")
		t.Error(id)
	}
}

// func TestWrite(t *testing.T) {
// 	tWriteFile := func(filename string, data []byte, perm os.FileMode) error {
// 		if filename != tId {
// 			return errors.New("filepath wrong, got: " + filename)
// 		}
// 		if perm != 0644 {
// 			return errors.New("perms wrong")
// 		}
// 		return nil
// 	}
// }
