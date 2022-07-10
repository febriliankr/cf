package service

import (
	"log"
	"net/http"

	"github.com/febriliankr/go-cfstore-api/internal/entities"

	"github.com/febriliankr/go-cfstore-api/internal/services"
	"github.com/labstack/echo/v4"
)

func (s *KantinService) UploadFileHandler(c echo.Context) error {

	appName := s.config.Server.UserTokenName

	form, err := c.MultipartForm()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	files := form.File["file"]

	key := services.GetSecretKeys()

	inUpload := entities.UploadMultipleFilesRequest{
		AppName: appName,
		Files:   files,
		Key:     key,
	}

	s3Urls, err := s.uc.UploadMultipleFiles(inUpload)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	res := entities.UploadMultipleFilesResponse{
		FilesURL: s3Urls,
	}

	return ResponseOK(c, http.StatusCreated, res)
}
