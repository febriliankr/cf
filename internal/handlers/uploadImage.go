package handler

import (
	"net/http"

	"github.com/febriliankr/go-cfstore-api/internal/entities"

	"github.com/febriliankr/go-cfstore-api/internal/response"
	"github.com/febriliankr/go-cfstore-api/internal/services"
	"github.com/febriliankr/go-cfstore-api/internal/usecases"
	"github.com/labstack/echo/v4"
)

func UploadFilesPublic(c echo.Context) error {

	appName := c.Param("app_name")

	if appName == "" {
		appName = "pldui"
	}

	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["file"]

	key := services.GetSecretKeys()

	inUpload := entities.UploadMultipleFilesRequest{
		AppName: appName,
		Files:   files,
		Key:     key,
	}

	s3Urls, err := usecases.UploadMultipleFiles(inUpload)

	if err != nil {
		return err
	}

	res := entities.UploadMultipleFilesResponse{
		FilesURL: s3Urls,
	}

	msg := "file uploaded successfully!"
	return response.SendResponse(http.StatusOK, c, response.Success(msg, res))
}
