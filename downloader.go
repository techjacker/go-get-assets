package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Downloader struct {
	OutputDir    string
	RelativePath string
	Assets       map[string]struct{}
}

func (d *Downloader) CreateDestPath(url string) (string, error) {

	idReg := regexp.MustCompile(`https://drive.google.com/file/d/(\w+)/.*`)
	id := idReg.FindStringSubmatch(url)

	if len(id) < 2 || id[1] == "" {
		return "", fmt.Errorf("%s", "no id found in url: "+url)
	}

	return "https://googledrive.com/host/" + id[1], nil
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
