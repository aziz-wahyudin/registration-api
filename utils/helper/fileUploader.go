package helper

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
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

func uploadFile(c *http.Request, keyPrefix string) (string, error) {

	file, fileheader, err := c.FormFile("file")
	if err != nil {
		log.Print(err)
		return "", err
	}
	defer file.Close()

	str := FileName(20)

	godotenv.Load(".env")

	s3Config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),                // bucket's name
		Key:         aws.String(keyPrefix + str + "-" + fileheader.Filename), // files destination location
		Body:        file,                                                    // content of the file
		ContentType: aws.String(fileheader.Header.Get("Content-Type")),       // content type
		Metadata: map[string]*string{
			"Original-Name": aws.String(fileheader.Filename),
		},
	}
	res, err := uploader.UploadWithContext(context.Background(), input)

	// return url location in aws
	return res.Location, err
}

func uploadPicture(w http.ResponseWriter, r *http.Request) {
	url, err := uploadFile(r, "picture/")
	if err != nil {
		http.Error(w, "Error uploading file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(url))
}
