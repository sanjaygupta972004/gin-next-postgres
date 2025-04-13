package config

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func GetAWSConfig(awsCfg *AWSConfig) (aws.Config, error) {
	if awsCfg == nil || awsCfg.Region == "" {
		return aws.Config{}, errors.New("AWS region is missing")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsCfg.Region),
	)
	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}

func isAWSConfigured(awsCfg *AWSConfig) bool {
	return awsCfg != nil && awsCfg.Region != "" && awsCfg.SesSenderEmail != "" && awsCfg.BucketName != ""
}

func NewSESClient() (*ses.Client, error) {
	awsCfg := GetGlobalConfig().AWS

	if !isAWSConfigured(&awsCfg) {
		return nil, errors.New("invalid AWS SES configuration")
	}

	cfg, err := GetAWSConfig(&awsCfg)
	if err != nil {
		return nil, err
	}

	return ses.NewFromConfig(cfg), nil
}

func NewS3Client() (*s3.Client, error) {
	awsCfg := GetGlobalConfig().AWS

	if !isAWSConfigured(&awsCfg) {
		return nil, errors.New("invalid AWS S3 configuration")
	}

	cfg, err := GetAWSConfig(&awsCfg)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}
