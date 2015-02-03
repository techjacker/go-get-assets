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

type Drive struct {
	*drive.Service
	client *http.Client
}

type Downloader struct {
	Assets    map[string]struct{}
	OutputDir string
	*Drive
}

func (d *Downloader) Multi() error {
	fmt.Println("downloader")
	return nil
}

// downloadAssets(assetList array, assetTypeLocation map)
func (d *Downloader) Single() error {
	return nil
}
