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
	var err error
	idReg := regexp.MustCompile(`https://drive.google.com/file/d/(\w+)/.*`)
	id := idReg.FindStringSubmatch(url)

	if id[1] == "" {
		err = fmt.Errorf("%s", "no id found in url: "+url)
	}

	return "https://googledrive.com/host/" + id[1], err
}

type writeFile func(filename string, data []byte, perm os.FileMode) error

func (d *Downloader) Download(url string, wFile writeFile) error {

	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	// err, destPath := d.CreateDestPath(url)
	// return wFile(destPath, data, 0644)

	return wFile(url, data, 0644)
}
