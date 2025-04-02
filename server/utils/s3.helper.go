package utils

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofrs/uuid"
)

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader, bucketName string, s3Client *s3.Client) (string, error) {
	fileExt := filepath.Ext(fileHeader.Filename)
	uid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %w", err)
	}
	fileName := fmt.Sprintf("%s%s", uid.String(), fileExt)

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buffer.Bytes()),
		ACL:    "public-read",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	s3URL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName)
	return s3URL, nil
}
