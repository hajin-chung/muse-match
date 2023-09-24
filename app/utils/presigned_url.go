package utils

import (
	"fmt"
	"musematch/app/globals"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3

func InitS3() error {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", globals.Env.ACCOUNT_ID)
	cred := credentials.NewStaticCredentials(
		globals.Env.ACCESS_KEY,
		globals.Env.SECRET_KEY,
		"",
	)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("auto"),
		Endpoint:    aws.String(endpoint),
		Credentials: cred,
	})
	if err != nil {
		return err
	}

	s3Client = s3.New(sess)
	return nil
}

func GetPutUrl(id string) (string, error) {
	req, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(globals.Env.BUCKET_NAME),
		Key:    aws.String(id),
	})
	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}
