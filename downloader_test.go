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

func TestCreateFilePath(t *testing.T) {
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

	err := d.Download(ts.URL, tWriteFile)
	if err != nil {
		t.Fatal(err)
	}
	// greeting, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Logf("%s", greeting)
}

func TestWriteToDisk(t *testing.T) {
	d := Downloader{
		os.TempDir() + "/go-get-assets",
		map[string]struct{}{
			"http://gdrive.com/traffic.jpg": {},
		},
	}
	t.Log(d)
}
