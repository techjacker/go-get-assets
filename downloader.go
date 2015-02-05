package main

import (
	"code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
	// "github.com/google/google-api-go-client/drive/v2"
	"net/http"
)

func NewDownloader(assets map[string]struct{}, outputDir string) (*Downloader, error) {
	// client, err := auth.GetOauth2Client(config.ClientId, config.ClientSecret, tokenPath, promptUser)

	client := http.Client{}

	drive, err := drive.New(&client)
	if err != nil {
		return nil, err
	}

	// Return a new authorized Drive client.
	return &Downloader{
		assets,
		outputDir,
		&Drive{
			drive,
			&client,
		},
	}, nil
	// return &Drive{drive, &client}, nil
}

type DFile struct {
	DownloadUrl   string `json:"downloadUrl,omitempty"`
	FileExtension string `json:"fileExtension,omitempty"`
	Title         string `json:"title,omitempty"`
	// EmbedLink string `json:"embedLink,omitempty"`
	// ModifiedDate string `json:"modifiedDate,omitempty"`
}

type Drive struct {
	*drive.Service
	client *http.Client
}

type Downloader struct {
	Assets    map[string]struct{}
	OutputDir string
	*Drive
}

func (d *Downloader) GetInfoAll() error {
	// map["http://gdrive.com/traffic.jpg":{}]
	for i, v := range d.Assets {
		v, err := d.GetInfo(i)
	}
	return err
}

// /home/andy/go/src/code.google.com/p/google-api-go-client/drive/v2/drive-gen.go
// line 3536
func (d *Downloader) GetInfo(url string) (*DFile, error) {
	return d.Service.Files.Get(url).Do()
}
