package helper

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// generate function to create random file name

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func nameGenerator(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func FileName(length int) string {
	return nameGenerator(length, charset)
}

//uploader

func UploadFile(r *http.Request, keyName string) (string, error) {
	file, fileHeader, err := r.FormFile(keyName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	str := FileName(20)
	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKey := os.Getenv("ACCESS_KEY_IAM")
	awsSecretKey := os.Getenv("SECRET_KEY_IAM")
	awsBucketName := os.Getenv("AWS_BUCKET_NAME")

	s3Config := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(awsBucketName),                                      // bucket's name
		Key:         aws.String("foto-hospital/" + str + "-" + fileHeader.Filename), // files destination location
		Body:        file,                                                           // content of the file
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),              // content type
	}
	res, err := uploader.UploadWithContext(r.Context(), input)
	if err != nil {
		return "", err
	}

	// return url location in aws
	return res.Location, nil
}
