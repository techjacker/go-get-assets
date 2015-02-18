package main

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {

	d := Downloader{
		os.TempDir() + "/go-get-assets",
		map[string]struct{}{
			"http://gdrive.com/traffic.jpg": {},
		},
	}

	d.Download("http://google.com")

	// if err != nil {
	// 	t.Fatal("error initialising drive", err)
	// }

	// t.Log(d)
}
