package entities

import (
	"mime/multipart"

	"github.com/febriliankr/go-cfstore-api/internal/services"
)

type (
	UploadMultipleFilesRequest struct {
		Files   []*multipart.FileHeader
		Key     services.AwsKeys
		AppName string
	}
	UploadMultipleFilesResponse struct {
		FilesURL []string `json:"files"`
	}
)
