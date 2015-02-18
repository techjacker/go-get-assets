package main

import (
	"fmt"
	"net/http"
)

type Downloader struct {
	OutputDir string
	Assets    map[string]struct{}
}

func (d *Downloader) Download(url string) error {

	c := new(http.Client)

	fmt.Print(c)
	fmt.Print(url)

	// if f.DownloadUrl == "" {
	// 	return DFile{}, fmt.Errorf("An error occurred: File is not downloadable")
	// }
	// if err != nil {
	// 	return fmt.Errorf("An error occurred: %v\n", err)
	// }
	return nil
}
