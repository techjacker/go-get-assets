package main

import (
	"io/ioutil"
	"net/http"
)

type Downloader struct {
	OutputDir string
	Assets    map[string]struct{}
}

func (d *Downloader) Download(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
