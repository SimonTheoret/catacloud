package client

// Each test must acquire the mutex before running. It ensures a consistent state for the test directory.

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	TestFileName = "some_file.txt"
	TestFilePath = "./tests_data/some_file.txt"
	// This mutex is essential for running the tests
	TestMutex = sync.Mutex{}
)

func setupAWSClient() AWSClient {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	},
	)
	bucket := os.Getenv("ARN_BUCKET")
	if err != nil {
		panic(err)
	}
	svc := s3.New(session)
	client := AWSClient{svc, bucket}
	return client
}

func TestListFileInTestBucket(t *testing.T) {
	TestMutex.Lock()
	defer TestMutex.Unlock()
	client := setupAWSClient()
	fileName := "list_dir/list_file_in_test_bucket_file.txt"
	err := client.UploadFile(TestFilePath, &fileName)
	if err != nil {
		panic(err)
	}
	files, err := client.ListFiles()
	if err != nil {
		panic(err)
	}
	actualLen := len(files.Files)
	expectedLen := 1
	if actualLen < expectedLen {
		msg := fmt.Errorf("expected len of at least %v, instead got len of %v", expectedLen, actualLen)
		panic(msg)
	}
}

func TestDeleteRemoteFile(t *testing.T) {
	TestMutex.Lock()
	defer TestMutex.Unlock()
	client := setupAWSClient()
	err := client.DeleteRemoteFile(TestFileName)
	if err != nil {
		panic(err)
	}
	remoteFiles, err := client.ListFiles()
	if err != nil {
		panic(err)
	}
	for _, file := range remoteFiles.Files {
		if file.Name == TestFileName {
			msg := fmt.Errorf("found unexpected file %v ", file.Name)
			panic(msg)
		}
	}

	err = client.UploadFile(TestFilePath, &TestFileName)
	if err != nil {
		panic(err)
	}
}

func TestUploadFile(t *testing.T) {
	TestMutex.Lock()
	defer TestMutex.Unlock()
	client := setupAWSClient()
	err := client.UploadFile(TestFilePath, &TestFilePath)
	if err != nil {
		panic(err)
	}
	files, err := client.ListFiles()
	if err != nil {
		panic(err)
	}
	notOK := true
	for _, file := range files.Files {
		if file.Name == TestFileName {
			notOK = false
		}
	}
	if notOK {
		panic(fmt.Errorf("did not find the uploaded file %v in the list of files %v", TestFileName, files))
	}
}
