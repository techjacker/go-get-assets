package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

type Downloader struct {
	OutputDir string
	Assets    map[string]struct{}
}

func (d *Downloader) CreateFilePath(url string) string {
	return "filePath"
}

type writeFile func(filename string, data []byte, perm os.FileMode) error

func (d *Downloader) Download(url string, wFile writeFile) error {

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

	// filePath := d.CreateFilePath(url)

	// return nil
	return wFile(url, data, 0644)
	// return wFile(filePath, data, 0644)
	// return ioutil.WriteFile(filePath, data, 0644)
}
