package awsput

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

func Put() {

	bucket := flag.String("b", "", "The bucket to upload the file to")
	filename := flag.String("f", "", "The file to upload")

	flag.Parse()

	if *bucket == "" || *filename == "" {
		fmt.Println("You must supply a bucket name (-b BUCKET) and file name (-f FILE)")
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(*filename)

	if err != nil {
		fmt.Println("Unable to open file " + *filename)
		return
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: bucket,
		Key:    filename,
		Body:   file,
	}

	_, err = PutFile(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got error uploading file:")
		fmt.Println(err)
		return
	}
}
