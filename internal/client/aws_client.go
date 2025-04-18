package client

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// AWSClient Wrapper around the AWS internal client. Its main function is to
// send requests and receive response from the S3(s).
type AWSClient struct {
	inner  *s3.S3
	bucket string
	// TODO: ADD CONFIG HERE?
}

func (c *AWSClient) ListFiles() (ListedFiles[s3.Object], error) {
	input := s3.ListObjectsV2Input{Bucket: &c.bucket}
	pages := make([]s3.ListObjectsV2Output, 0)
	PagesAccumulator := func(val *s3.ListObjectsV2Output, _ bool) bool {
		pages = append(pages, *val)
		return true
	}
	err := c.inner.ListObjectsV2Pages(&input, PagesAccumulator)
	if err != nil {
		var ret ListedFiles[s3.Object]
		return ret, err
	}
	files, objects := c.populateListOFFiles(pages)
	listedFiles := ListedFiles[s3.Object]{Inner: objects, Files: files}

	return listedFiles, nil
}

// populateListOFFiles Iterate over the given pages and gives back a list of
// files with metadata.
func (c *AWSClient) populateListOFFiles(pages []s3.ListObjectsV2Output) ([]FileWithMetadata, []s3.Object) {
	listOfFiles := make([]FileWithMetadata, 0)
	listOfObjects := make([]s3.Object, 0)
	for i := range pages {
		elem := pages[i]
		for j := range elem.Contents {
			awsFile := elem.Contents[j]
			// TODO:: Use a logger instead of a raw Println
			fmt.Println(awsFile)
			file := FileWithMetadata{LastModified: *awsFile.LastModified, Name: *awsFile.Key}
			listOfFiles = append(listOfFiles, file)
			listOfObjects = append(listOfObjects, *awsFile)
		}

	}
	return listOfFiles, listOfObjects
}

func (c *AWSClient) DeleteRemoteFile(file string) error {
	_, err := c.inner.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(c.bucket), Key: aws.String(file)})
	if err != nil {
		return err
	}

	err = c.inner.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(file),
	})
	return err
}

func (c *AWSClient) UploadFile(filepath string, key *string) error {
	uploader := s3manager.NewUploaderWithClient(c.inner)
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	uploadInput := s3manager.UploadInput{
		Bucket: &c.bucket,
		Key:    key,
		Body:   file,
	}
	_, err = uploader.Upload(&uploadInput)
	if err != nil {
		return err
	}
	return file.Close()
}
