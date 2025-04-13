package utils

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofrs/uuid"
	"github.com/savvy-bit/gin-react-postgres/config"
)

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileExt := filepath.Ext(fileHeader.Filename)
	uid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %w", err)
	}

	awsConfig := config.GetGlobalConfig().AWS
	bucketName := awsConfig.BucketName
	region := awsConfig.Region
	if bucketName == "" || region == "" {
		return "", fmt.Errorf("AWS configuration is not valid")
	}

	s3Client, err := config.NewS3Client()
	if err != nil {
		return "", fmt.Errorf("failed to create S3 client: %w", err)
	}

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	defer file.Close()

	contentType := http.DetectContentType(buffer.Bytes()[:512])
	if !strings.HasPrefix(contentType, "image/") {
		return "", fmt.Errorf("invalid file type: %s. Only images are allowed", contentType)
	}

	fileName := fmt.Sprintf("%s%s", uid.String(), fileExt)

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: aws.String(contentType),
		ACL:         "public-read",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	s3URL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, fileName)
	return s3URL, nil
}
