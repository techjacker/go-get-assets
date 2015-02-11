package main

import (
	"code.google.com/p/google-api-go-client/drive/v2"
	// "github.com/google/google-api-go-client/drive/v2"
	"net/http"
)

func ExtractId(url string) string {
	return ""
}

func NewDownloader(assets map[string]struct{}, outputDir string) (*Downloader, error) {

	client := http.Client{}
	drive, err := drive.New(&client)

	if err != nil {
		return nil, err
	}

	return &Downloader{
		outputDir,
		assets,
		map[string]DFile{},
		&Drive{
			drive.Files,
			&client,
		},
	}, nil
}

type Downloader struct {
	OutputDir string
	Assets    map[string]struct{}
	Metadata  map[string]DFile
	Drive     *Drive
}

func (d *Downloader) GetInfoAll() error {
	// map["http://gdrive.com/traffic.jpg":{}]
	var err error
	for i := range d.Assets {
		d.Metadata[i], err = d.GetInfo(i, d.Drive.DFilesService)
	}
	return err
}

// /home/andy/go/src/code.google.com/p/google-api-go-client/drive/v2/drive-gen.go
// line 3536
func (d *Downloader) GetInfo(id string, dService DFilesService) (DFile, error) {
	f, err := dService.Get(id).Do()
	return DFile{
		f.DownloadUrl,
		f.FileExtension,
		f.Title,
	}, err
}

func (d *Downloader) Download(url string) error {
	// if f.DownloadUrl == "" {
	// 	return DFile{}, fmt.Errorf("An error occurred: File is not downloadable")
	// }
	// if err != nil {
	// 	return fmt.Errorf("An error occurred: %v\n", err)
	// }
	return nil
}
