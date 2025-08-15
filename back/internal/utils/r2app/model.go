package r2app

import "github.com/aws/aws-sdk-go-v2/service/s3"

type R2Client struct {
    Bucket    string
    AccountID string
    Client    *s3.Client
}