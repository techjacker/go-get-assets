package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type writeFile func(filename string, data []byte, perm os.FileMode) error

type Downloader struct {
	OutputDir    string
	RelativePath string
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

func (d *Downloader) Download(url string, chanDown chan Res) {
	res, err := http.Get(url)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	chanDown <- Res{data, err}
}

func (d *Downloader) WriteToDisk(destPath string, chanDown chan Res) {
	res := <-chanDown
	if res.Err != nil {
		ioutil.WriteFile(destPath, res.Data, 0644)
	}
}

func (d *Downloader) Run(assets map[string]Asset) error {
	chanDown := make(chan Res)
	for k, _ := range assets {
		if id := d.ExtractId(k); id != "" {
			go d.Download(d.CreateTargetUrl(id), chanDown)
			go d.WriteToDisk(d.CreateDestPath(id), chanDown)
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
