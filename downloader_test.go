package main

import (
	"os"
	"testing"
)

func TestNewDownloader(t *testing.T) {
	assets := map[string]struct{}{
		"http://gdrive.com/traffic.jpg": {},
	}
	var out string = os.TempDir() + "/go-get-assets"

	d, err := NewDownloader(assets, out)

	if err != nil {
		t.Fatal("error initialising drive", err)
	}

	// t.Logf("%q", d.Drive.DFilesService)
	t.Logf("%q", d.Drive)
	// t.Log(d)
}
