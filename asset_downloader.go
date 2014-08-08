package main

import (
	"fmt"
)

type AssetDownloader struct {
	assetUrls []string
	outputDir string
}

func (*AssetDownloader) downloadAssets(assetUrls []string, outputDir string) error {
	fmt.Println("asset downloader")
	return nil
}

// downloadAssets(assetList array, assetTypeLocation map)
func (*AssetDownloader) downloadSingleAsset(assetUrl string, outputDir string) error {
	return nil
}
