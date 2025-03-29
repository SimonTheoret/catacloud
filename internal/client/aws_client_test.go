package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type testingConfig struct {
	bucket string
}

func TestListFileInTestBucket(t *testing.T) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	},
	)
	bucket := os.Getenv("ARN_BUCKET")
	if err != nil {
		t.Fatal(err)
	}
	svc := s3.New(session)
	client := AWSClient{svc, bucket}
	files := client.ListFiles()
	actualLen := len(files.Files)
	expectedLen := 1
	if actualLen != expectedLen {
		msg := fmt.Errorf("Expected len of %v, instead got len of %v", expectedLen, actualLen)
		panic(msg)
	}
}
