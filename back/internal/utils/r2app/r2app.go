package r2app

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewR2Client(bucketName, accountId, accessKeyId, accessKeySecret string) *s3.Client {
cfg, err := config.LoadDefaultConfig(context.TODO(),
    config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
    config.WithRegion("auto"),
  )
  if err != nil {
    log.Fatal(err)
  }

  client := s3.NewFromConfig(cfg, func(o *s3.Options) {
    o.EndpointResolver = s3.EndpointResolverFromURL(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
  })

	return client
}


func ListObjects(client *s3.Client, bucketName string) {
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range output.Contents {
		fmt.Printf("%+v\n", object)
	}
}

func ListBuckets(client *s3.Client) {
	output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range output.Buckets {
		fmt.Printf("%+v\n", bucket)
	}
}
