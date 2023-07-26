package s3

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/jhyehuang/crontab_server/src/configs"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
)

func UnloadAwsS3(file io.Reader, fileName string) (string, error) {

	awsConfig, _ := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(configs.GetValue(configs.AwsAccessKey, configs.Get().AWS.AwsAccessKeyId), configs.GetValue(configs.AwsSecretKey, configs.Get().AWS.AwsSecretAccessKey), ""),
		),
		config.WithRegion(configs.GetValue(configs.AwsRegion, configs.Get().AWS.AwsDefaultRegion)),
	)

	client := s3.NewFromConfig(awsConfig)
	bucket := configs.GetValue(configs.AwsStaticBucket, configs.Get().AWS.StaticResBucket)

	readerForS3, bufW := io.Pipe()
	defer bufW.Close()
	readerForMD5 := io.TeeReader(file, bufW)

	h := md5.New()
	io.Copy(h, readerForS3)
	fileKey := base64.URLEncoding.EncodeToString(h.Sum(nil)) + fileName

	putObjectArgs := s3.PutObjectInput{
		ACL:    types.ObjectCannedACLPublicRead,
		Body:   readerForMD5,
		Bucket: &bucket,
		Key:    &fileKey,
	}
	res, err := client.PutObject(
		context.Background(),
		&putObjectArgs)
	if err != nil {
		log.Errorf("failed to presign request, %v\n", err)
		return "", err
	}
	log.Infof("%+v", res)

	// 判断是否上传成功
	if res.ETag == nil {
		return "", fmt.Errorf("upload error")
	}

	// 拼接返回的url
	url := fmt.Sprintf("%s/%s", "http://"+bucket+".s3.cn-north-1.amazonaws.com.cn", fileKey)

	// 测试一下url是否可用 返回200
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("url is not available")
	}

	return url, nil

}

func PresignAwsS3(fileName string) (string, *http.Header, error) {

	awsConfig, _ := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(configs.GetValue(configs.AwsAccessKey, configs.Get().AWS.AwsAccessKeyId), configs.GetValue(configs.AwsSecretKey, configs.Get().AWS.AwsSecretAccessKey), ""),
		),
		config.WithRegion(configs.GetValue(configs.AwsRegion, configs.Get().AWS.AwsDefaultRegion)),
	)

	client := s3.NewFromConfig(awsConfig)

	bucket := configs.GetValue(configs.AwsStaticBucket, configs.Get().AWS.StaticResBucket)

	presignClient := s3.NewPresignClient(client)
	expiration := time.Hour * 24

	putObjectArgs := s3.PutObjectInput{
		ACL:    types.ObjectCannedACLPublicRead,
		Bucket: &bucket,
		Key:    &fileName,
	}
	res, err := presignClient.PresignPutObject(
		context.Background(),
		&putObjectArgs,
		s3.WithPresignExpires(expiration),
	)
	if err != nil {
		log.Errorf("failed to presign request, %v\n", err)
		return "", nil, err
	}
	log.Infof("%+v", res.URL)
	log.Infof("%+v", res.SignedHeader)

	// 判断是否成功
	if res.URL == "" {
		return "", nil, fmt.Errorf("upload error")
	}

	return res.URL, &res.SignedHeader, nil

}
