package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	var d Downloader
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := d.Download(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("res", res)
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
