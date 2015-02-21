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

var tId string = "0ByPfUp1fLihSNm5SSjZoalhPQ3M"
var tIn string = "https://drive.google.com/file/d/" + tId + "/view?usp=sharing"

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
	d.OutputDir = "src/images"
	got := d.CreateDestPath(tIn)
	if got != "src/images/"+tId {
		t.Fatal("got:", got)
	}
}

func TestCreateTargetUrl(t *testing.T) {
	var d Downloader
	want := "https://googledrive.com/host/" + tId
	if d.CreateTargetUrl(tIn) != want {
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
		if filename != ts.URL {
			return errors.New("filepath wrong")
		}
		if strings.TrimSpace(string(data)) != "Hello, client" {
			return errors.New("data wrong")
		}
		if perm != 0644 {
			return errors.New("perms wrong")
		}
		return nil
	}

	tCreatePath := func(s string) (string, error) {
		return s, nil
	}

	err := d.Download(ts.URL, tWriteFile, tCreatePath)

	if err != nil {
		t.Fatal(err)
	}
}

// func TestWriteToDisk(t *testing.T) {
// 	d := Downloader{
// 		os.TempDir() + "/go-get-assets/download",
// 		"/images",
// 		map[string]struct{}{
// 			"http://gdrive.com/traffic.jpg": {},
// 		},
// 	}
// 	t.Log(d)
// }
