package r2app

import (
	"log"
	"testing"

	"github.com/nsevenpack/env/env"
)

func TestNewR2Client_success(t *testing.T) {
	bucketName := env.Get("R2_BUCKET_NAME")
	accountId := env.Get("R2_ACCOUNT_ID")
	accessKeyId := env.Get("R2_ACCESS_KEY_ID")
	accessKeySecret := env.Get("R2_ACCESS_KEY_SECRET")

	client := NewR2Client(bucketName, accountId, accessKeyId, accessKeySecret)

	if client == nil {
		t.Fatal("Echec de la creation du client R2 : client est nil")
	}

	log.Printf("Client R2 cree avec succes : %#v\n", client)
}

func TestListObjects_success(t *testing.T) {
	bucketName := env.Get("R2_BUCKET_NAME")
	accountId := env.Get("R2_ACCOUNT_ID")
	accessKeyId := env.Get("R2_ACCESS_KEY_ID")
	accessKeySecret := env.Get("R2_ACCESS_KEY_SECRET")

	client := NewR2Client(bucketName, accountId, accessKeyId, accessKeySecret)

	if client == nil {
		t.Fatal("Echec de la creation du client R2 : client est nil")
	}

	ListObjects(client, bucketName)
}

func TestListBuckets_success(t *testing.T) {
	bucketName := env.Get("R2_BUCKET_NAME")
	accountId := env.Get("R2_ACCOUNT_ID")
	accessKeyId := env.Get("R2_ACCESS_KEY_ID")
	accessKeySecret := env.Get("R2_ACCESS_KEY_SECRET")

	client := NewR2Client(bucketName, accountId, accessKeyId, accessKeySecret)

	if client == nil {
		t.Fatal("Echec de la creation du client R2 : client est nil")
	}

	ListBuckets(client)
}