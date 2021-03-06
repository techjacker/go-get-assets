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
	tID      = "0ByPfUp1fLihSNm5SSjZoalhPQ3M"
	tIn      = gDriveURL + "/" + tID + "/view?usp=sharing"
	tRel     = "/"
	eDir     = "src/images/"
	eRelPath = tRel + tID
	ePath    = eDir + tID
	eURL     = "https://googledrive.com/host/" + tID
)

// func TestExtractIDMissing(t *testing.T) {
// bad string goes here

func TestExtractID(t *testing.T) {
	d := NewDownloader("", "")
	got := d.Extracter.Gdrive(tIn)
	if got != tID {
		t.Fatal("got:", got)
	}
}

func TestCreateDestPath(t *testing.T) {
	var d Downloader
	d.OutputDir = eDir
	got := d.CreateDestPath(tID)
	if got != ePath {
		t.Fatal("got:", got)
	}
}

func TestCreateRelPath(t *testing.T) {
	var d Downloader
	d.RelativePath = tRel
	if d.CreateRelPath(tID) != eRelPath {
		t.Fatal("got:", d.CreateRelPath(tID))
	}
}

func TestCreateTargetUrl(t *testing.T) {
	var d Downloader
	if d.CreateTargetUrl(tID) != eURL {
		t.Fatal("got:", d.CreateTargetUrl(tID))
	}
}

type AssetsTest struct {
	Mapofphotos struct {
		Photourl  string
		Somearray []string
		Notaphoto string
	}
}

func TestDownload(t *testing.T) {
	var (
		d        Downloader
		res      Res
		body     = "hello"
		id       = "12345"
		origURL  = "http://sdffdsfsd.com"
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
	go d.Download(ts.URL, origURL, id, chanDown)
	<-done // wait for output to be received by res

	if res.Err != nil {
		t.Error("error is not nil")
		t.Error(res.Err)
	}
	if strings.TrimSpace(string(res.Data)) != body {
		t.Error("data wrong")
		t.Error(res.Data)
	}
	if res.Url != origURL {
		t.Error("origURL not passed")
		t.Error(origURL)
	}
	if res.Id != id {
		t.Error("id not passed")
		t.Error(id)
	}
}

// func TestWrite(t *testing.T) {
// 	tWriteFile := func(filename string, data []byte, perm os.FileMode) error {
// 		if filename != tID {
// 			return errors.New("filepath wrong, got: " + filename)
// 		}
// 		if perm != 0644 {
// 			return errors.New("perms wrong")
// 		}
// 		return nil
// 	}
// }
