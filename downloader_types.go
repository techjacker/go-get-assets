package main

import (
	"code.google.com/p/google-api-go-client/drive/v2"
	// "github.com/google/google-api-go-client/drive/v2"
	// "net/http"
)

type DFile struct {
	DownloadUrl   string `json:"downloadUrl,omitempty"`
	FileExtension string `json:"fileExtension,omitempty"`
	Title         string `json:"title,omitempty"`
	// EmbedLink string `json:"embedLink,omitempty"`
	// ModifiedDate string `json:"modifiedDate,omitempty"`
}

type DGetCall interface {
	Do() (*drive.File, error)
	// Do() (*DFile, error)
}

type DFilesService interface {
	// Get(string) *struct{}
	// Get(string) *interface{}
	// Get(string) *interface {
	// 	Do() (*drive.File, error)
	// }
	// Get(string) *DGetCall
	Get(string) *drive.FilesGetCall
}

// type Drive struct {
// 	Service DFilesService
// 	client  *http.Client
// 	// *drive.Service
// }

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
