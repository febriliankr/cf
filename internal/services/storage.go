package services

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsKeys struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
}

func GetSecretKeys() AwsKeys {
	AWS_REGION := os.Getenv("AWS_REGION")
	AWS_ACCESS_KEY_ID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_KEY := os.Getenv("AWS_SECRET_KEY")
	BUCKET_NAME := os.Getenv("BUCKET_NAME")

	keys := AwsKeys{
		Region:          AWS_REGION,
		AccessKeyID:     AWS_ACCESS_KEY_ID,
		SecretAccessKey: AWS_SECRET_KEY,
		BucketName:      BUCKET_NAME,
	}

	return keys
}

func UploadFilePublic(session *session.Session, sourceFileDir string, filename string) (*s3.PutObjectOutput, error) {

	var out *s3.PutObjectOutput
	upFile, err := os.Open(sourceFileDir)
	if err != nil {
		return out, err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	bucketName := os.Getenv("BUCKET_NAME")

	out, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(filename),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return out, err
}
