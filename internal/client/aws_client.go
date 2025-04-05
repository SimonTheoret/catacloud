package client

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
)

// AWSClient Wrapper around the AWS internal client. Its main function is to
// send requests and receive response from the S3(s).
type AWSClient struct {
	inner  *s3.S3
	bucket string
	// TODO: Add config/pointer towards a config here?
}

func (c *AWSClient) ListFiles() ListedFiles[s3.Object] {
	input := s3.ListObjectsV2Input{Bucket: &c.bucket}
	pages := make([]s3.ListObjectsV2Output, 0)
	PagesAccumulator := func(val *s3.ListObjectsV2Output, _ bool) bool {
		pages = append(pages, *val)
		return true
	}
	err := c.inner.ListObjectsV2Pages(&input, PagesAccumulator)
	if err != nil {
		// TODO: We should handle errors more gracefully.
		// TODO: Log errors!
		panic(err)
	}
	files, objects := c.populateListOFFiles(pages)
	listedFiles := ListedFiles[s3.Object]{Inner: objects, Files: files}

	return listedFiles
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
			// TODO:: Use a logger instead of a raw Println ?
			fmt.Println(awsFile)
			file := FileWithMetadata{LastModified: *awsFile.LastModified, Name: *awsFile.Key}
			listOfFiles = append(listOfFiles, file)
			listOfObjects = append(listOfObjects, *awsFile)
		}

	}
	return listOfFiles, listOfObjects
}

func (c *AWSClient) DownloadFiles(files []string) (err error) {
	return fmt.Errorf("TODO")
}

func (c *AWSClient) UploadFiles(filepaths []string) (err error) {
	return fmt.Errorf("TODO")
}
