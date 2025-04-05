package email

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/savvy-bit/gin-react-postgres/config"
)

func isAWSConfigured(awsConfig *config.AWSConfig) bool {
	if awsConfig == nil {
		return false
	}
	return awsConfig.Region != "" || awsConfig.SesSenderEmail != "" || awsConfig.BucketName != ""
}

func GetAWSConfig(configVariables *config.AWSConfig) (aws.Config, error) {
	if configVariables == nil {
		return aws.Config{}, errors.New("AWS configuration is not valid")
	}

	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(configVariables.Region),
	)

	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}

func SESAWSClient() (*ses.Client, error) {
	awsConfig := config.GetGlobalConfig().AWS

	isValid := isAWSConfigured(&awsConfig)
	if !isValid {
		return nil, errors.New("AWS configuration is not valid")
	}

	cfg, err := GetAWSConfig(&awsConfig)
	if err != nil {
		return nil, err
	}
	return ses.NewFromConfig(cfg), nil
}

func S3AWSClient() (*s3.Client, error) {
	awsConfig := config.GetGlobalConfig().AWS

	isValid := isAWSConfigured(&awsConfig)
	if !isValid {
		return nil, errors.New("AWS configuration is not valid")
	}

	cfg, err := GetAWSConfig(&awsConfig)
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg), nil
}
