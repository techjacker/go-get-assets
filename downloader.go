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
	Assets       map[string]struct{}
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

func (d *Downloader) CreateDestPath(id string) string {
	return filepath.Join(d.OutputDir + id)
}
func (d *Downloader) CreateTargetUrl(id string) string {
	return "https://googledrive.com/host/" + id
}

type modifyStr func(s string) (string, error)
type writeFile func(filename string, data []byte, perm os.FileMode) error

func (d *Downloader) Download(url string, wFile writeFile, createPathStr modifyStr) error {

	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	destPath, err := createPathStr(url)

	if err != nil {
		return err
	}

	return wFile(destPath, data, 0644)
}
