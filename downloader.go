package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type Downloader struct {
	OutputDir    string
	RelativePath string
	ChanDown     chan Res
}

type Res struct {
	Data []byte
	Err  error
}

func (d *Downloader) ExtractId(url string) string {

	idReg := regexp.MustCompile(`https://drive.google.com/file/d/(\w+)/.*`)
	id := idReg.FindStringSubmatch(url)

	// didn't find a match
	if len(id) < 2 {
		return ""
	}
	return id[1]
}

func (d *Downloader) CreateRelPath(id string) string {
	return filepath.Join(d.RelativePath, id)
}
func (d *Downloader) CreateDestPath(id string) string {
	return filepath.Join(d.OutputDir, id)
}
func (d *Downloader) CreateTargetUrl(id string) string {
	return "https://googledrive.com/host/" + id
}

type writeFile func(filename string, data []byte, perm os.FileMode) error

func (d *Downloader) Download(id string, url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	d.ChanDown <- Res{data, err}
	return data, err
}

func (d *Downloader) Run(assets map[string]Asset) error {
	for k, _ := range assets {
		if id := d.ExtractId(k); id != "" {
			go d.Download(id, d.CreateTargetUrl(id))
		}
	}
	return nil
}

// ioutil.WriteFile

// if err != nil {
// 	d.Assets[k] = Asset{"", err}
// } else {
// 	d.Assets[k] = Asset{d.CreateRelPath(id), nil}
// }
