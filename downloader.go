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
	Assets       map[string]Asset
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

func (d *Downloader) Download(id string, url string, wFile writeFile) error {

	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return wFile(d.CreateDestPath(id), data, 0644)
}

type Res struct {
	res string
}

func (d *Downloader) Run() error {

	var (
		id  string
		err error
	)

	for k, _ := range d.Assets {
		if id = d.ExtractId(k); id != "" {
			err = d.Download(id, d.CreateTargetUrl(id), ioutil.WriteFile)
			if err != nil {
				d.Assets[k] = Asset{"", err}
			} else {
				d.Assets[k] = Asset{d.CreateRelPath(id), nil}
			}
		}
	}
	return err
}
