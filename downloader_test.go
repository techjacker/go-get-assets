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

func TestCreateDestPathMissing(t *testing.T) {

	// bad string goes here

}

func TestCreateDestPath(t *testing.T) {
	var d Downloader
	d.OutputDir = "src/images"
	d.RelativePath = "/images"

	id := "0ByPfUp1fLihSNm5SSjZoalhPQ3M"
	in := "https://drive.google.com/file/d/" + id + "/view?usp=sharing"
	got, err := d.CreateDestPath(in)
	want := "https://googledrive.com/host/" + id

	if got != want {
		t.Fatal("got:", got)
	}

	if err != nil {
		t.Fatal("err:", err)
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
