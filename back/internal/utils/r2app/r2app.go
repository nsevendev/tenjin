package r2app

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func R2Client(bucketName, accountId, accessKeyId, accessKeySecret string) *s3.Client{
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
