package main

import (
	// "code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
)

type Downloader struct {
	Assets    map[string]struct{}
	outputDir string
}

func (d *Downloader) downloadAssets() error {
	fmt.Println("downloader")
	return nil
}

// downloadAssets(assetList array, assetTypeLocation map)
func (d *Downloader) downloadSingleAsset() error {
	return nil
}
