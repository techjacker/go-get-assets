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

func TestDownload(t *testing.T) {
	var (
		d   Downloader
		res = "hello"
	)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, res)
	}))
	defer ts.Close()

	data, err := d.Download(tId, ts.URL)

	if err != nil {
		t.Error(err)
	}
	if strings.TrimSpace(string(data)) != res {
		t.Error("data wrong")
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
