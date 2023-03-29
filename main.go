package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	endpint    = os.Getenv("ALI_ANDPINT")
	acessKey   = os.Getenv("ALI_ACESS_KEY")
	secretKey  = os.Getenv("ALI_SECRET_KEY")
	bucketName = os.Getenv("ALI_BUCKET_NAME")
	UploadFile = ""
)

func Checkenv() error {
	if endpint == "" || acessKey == "" || secretKey == "" {
		return errors.New("请检查登录信息")
	}
	return nil
}

func loadParams() {
	flag.StringVar(&UploadFile, "f", "", "指定文件")
	flag.Parse()

}

func upload(filePath string) error {

	client, err := oss.New(endpint, acessKey, secretKey)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(filePath, filePath)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	loadParams()
	if err := Checkenv(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := upload(UploadFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
