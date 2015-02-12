package main

import (
	"code.google.com/p/google-api-go-client/drive/v2"
	"net/http"
)

type MyGetCall interface {
	Do() (*drive.File, error)
}

type MyFilesService interface {
	Get(string) *MyGetCall
	// Get(string) *drive.FilesGetCall
}

func main() {
	drive, _ := drive.New(new(http.Client))
	GetInfo("id", drive.Files)
}

func GetInfo(id string, fService MyFilesService) {
	fService.Get(id).Do()
}
