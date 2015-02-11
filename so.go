package main

import (
	"code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
	// "github.com/google/google-api-go-client/drive/v2"
	"net/http"
)

func DriveFactory() {
	client := http.Client{}
	drive, err := drive.New(&client)

	if err != nil {
		fmt.Errorf("%s", "error initing g drive")
	}
	GetInfo("id", drive.Files)
}

// /home/andy/go/src/code.google.com/p/google-api-go-client/drive/v2/drive-gen.go
// line 3536
func GetInfo(id string, fService MyFilesService) (*drive.File, error) {
	return fService.Get(id).Do()
}
