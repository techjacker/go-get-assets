package main

import (
	// "code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
)

type Downloader struct {
	assetUrls []string
	outputDir string
}

func (*Downloader) downloadAssets(assetUrls []string, outputDir string) error {
	fmt.Println("asset downloader")
	return nil
}

// downloadAssets(assetList array, assetTypeLocation map)
func (*Downloader) downloadSingleAsset(assetUrl string, outputDir string) error {
	return nil
}
