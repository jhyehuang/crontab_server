package s3

import (
	"fmt"
	"os"
	"testing"
)

func TestAwsS3(t *testing.T) {

	fileKey := "awsS3.go"

	// 打开文件 fileKey
	file, err := os.Open(fileKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	res, err := UnloadAwsS3(file, fileKey)
	if err != nil {
		fmt.Printf("failed to presign request, %v\n", err)
		return
	}
	fmt.Printf("%s \n", res)
}
