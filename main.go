package main

import (
	"context"

	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// BucketBasics encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.

// DownloadFile gets an object from a bucket and stores it in a local file.
func DownloadFile(bucketName string, objectKey string) ([]byte, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("LoadDefaultConfig:", err)
		return []byte{}, err
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	result, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("NewFromConfig:", err)
		return []byte{}, err
	}
	defer result.Body.Close()

	return ioutil.ReadAll(result.Body)

}

func handler(ctx context.Context, s3Event events.S3Event) error {
	log.Println("env1", os.Getenv("AWS_ACCESS_KEY_ID"))
	log.Println("env2", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	log.Println("env3", os.Getenv("AWS_REGION"))
	//log.Printf("%#v", s3Event)
	for _, li := range s3Event.Records {
		log.Println("get file ->:", li.S3.Object.Key)
		data, err := DownloadFile(li.S3.Bucket.Name, li.S3.Object.Key)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(data))
		}

	}
	return nil
}

func main() {
	lambda.Start(handler)
}
