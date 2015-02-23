package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	tId   = "0ByPfUp1fLihSNm5SSjZoalhPQ3M"
	eDir  = "src/images/"
	ePath = eDir + tId
	eUrl  = "https://googledrive.com/host/" + tId
	tIn   = "https://drive.google.com/file/d/" + tId + "/view?usp=sharing"
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

func TestCreateTargetUrl(t *testing.T) {
	var d Downloader
	if d.CreateTargetUrl(tId) != eUrl {
		t.Fatal("got:", d.CreateTargetUrl(tIn))
	}
}

func TestDownload(t *testing.T) {
	var d Downloader
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	tWriteFile := func(filename string, data []byte, perm os.FileMode) error {
		if filename != tId {
			return errors.New("filepath wrong, got: " + filename)
		}
		if strings.TrimSpace(string(data)) != "Hello, client" {
			return errors.New("data wrong")
		}
		if perm != 0644 {
			return errors.New("perms wrong")
		}
		return nil
	}

	err := d.Download(tId, ts.URL, tWriteFile)

	if err != nil {
		t.Fatal(err)
	}
}
