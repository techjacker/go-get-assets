package main

import (
	"io/ioutil"
	"net/http"
)

type Downloader struct {
	OutputDir string
	Assets    map[string]struct{}
}

func (d *Downloader) CreateFilePath(url string) string {
	return "filePath"
}

func (d *Downloader) Download(url string) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// return ioutil.ReadAll(res.Body)
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	filePath := d.CreateFilePath(url)

	return nil
	// return ioutil.WriteFile(filePath, data, 0644)
}
