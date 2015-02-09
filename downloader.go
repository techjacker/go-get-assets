package main

import (
	"code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
	// "fmt"
	// "github.com/google/google-api-go-client/drive/v2"
	"net/http"
)

func ExtractId(url string) string {
	return ""
}

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
			drive.Files,
			&client,
		},
		map[string]DFile{},
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

type DFilesService interface {
	// Get(string) *DGetCall
	Get(string) *drive.FilesGetCall
}

type DGetCall interface {
	// Do() (*DFilesService, error)
	Do() (*drive.File, error)
}

type Drive struct {
	DFilesService
	client *http.Client
	// *drive.Service
}

type Downloader struct {
	Assets    map[string]struct{}
	OutputDir string
	*Drive
	Metadata map[string]DFile
}

func (d *Downloader) GetInfoAll() error {
	// map["http://gdrive.com/traffic.jpg":{}]
	var err error
	for i := range d.Assets {
		d.Metadata[i], err = d.GetInfo(i)
	}
	return err
}

// /home/andy/go/src/code.google.com/p/google-api-go-client/drive/v2/drive-gen.go
// line 3536
func (d *Downloader) GetInfo(id string) (DFile, error) {

	fmt.Printf("q", d.Drive)
	// f, err := d.Service.Files.Get(id).Do()
	// f, err := d.Files.Get(id).Do()
	f, err := struct{}{}, fmt.Errorf("")

	fmt.Print(f)
	// if f.DownloadUrl == "" {
	// 	return DFile{}, fmt.Errorf("An error occurred: File is not downloadable")
	// }

	return DFile{}, err
	// return DFile{
	// 	f.DownloadUrl,
	// 	f.FileExtension,
	// 	f.Title,
	// }, err
}

func (d *Downloader) Download(url string) error {
	// if err != nil {
	// 	return fmt.Errorf("An error occurred: %v\n", err)
	// }
	return nil
}

// type FilesService struct {
// 	s *Service
// }
// func (r *FilesService) Get(fileId string) *FilesGetCall {
// type FilesGetCall struct {
// 	s      *Service
// 	fileId string
// 	opt_   map[string]interface{}
// }
// func (c *FilesGetCall) Do() (*File, error) {
