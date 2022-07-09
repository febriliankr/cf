package usecases

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/febriliankr/go-cfstore-api/internal/helper"
	"github.com/febriliankr/go-cfstore-api/internal/services"
)

func UploadMultipleFiles(in entities.UploadMultipleFilesRequest) ([]string, error) {

	var s3Urls []string

	for _, file := range in.Files {
		url, err := UploadFile(file, in.Key, in.AppName)
		if err != nil {
			return s3Urls, err
		}
		s3Urls = append(s3Urls, url)
	}

	return s3Urls, nil
}

func UploadFile(file *multipart.FileHeader, key services.AwsKeys, directory string) (string, error) {

	var s3Url string

	filename := helper.GenerateUUID() + "_" + file.Filename

	tmpFileDir := fmt.Sprintf("/tmp/%s", filename)

	source, err := file.Open()

	if err != nil {
		return s3Url, err
	}

	destination, err := os.Create(fmt.Sprintf("/tmp/%s", filename))
	if err != nil {
		return s3Url, err
	}

	if _, err = io.Copy(destination, source); err != nil {
		return s3Url, err
	}

	// Close the file without defer to close before next iteration of loop
	source.Close()
	destination.Close()

	// Upload to S3
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(key.Region)},
	)

	if err != nil {
		return s3Url, err
	}

	extension, _ := fileNameExtension(filename)
	s3FileDestination := directory + "/" + filename + extension

	_, err = services.UploadFilePublic(sess, tmpFileDir, s3FileDestination)

	if err != nil {
		return s3Url, err
	}

	s3Url = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", key.BucketName, key.Region, s3FileDestination)

	err = os.Remove(tmpFileDir)

	if err != nil {
		return s3Url, err
	}

	return s3Url, nil
}

func fileNameExtension(filename string) (string, string) {

	var extension = filepath.Ext(filename)
	var name = filename[0 : len(filename)-len(extension)]

	return extension, name
}
