package main

import (
	"code.google.com/p/google-api-go-client/drive/v2"
	"net/http"
)

type MyFile struct {
	DownloadUrl string `json:"downloadUrl,omitempty"`
}

type MyFilesGetCall interface {
	Do() (*MyFile, error)
	// Do() (*drive.File, error) // this DOES work
}

type MyFilesService interface {
	Get(string) *MyFilesGetCall
	// Get(string) *drive.FilesGetCall // this DOES work
}

func GetInfo(id string, myService MyFilesService) {
	myService.Get(id).Do()
}

func main() {
	drive, _ := drive.New(new(http.Client))
	GetInfo("id", drive.Files)
}
