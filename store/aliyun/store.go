package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/open-huajia/cloud-station-go/store"
	"os"
)

var (
	// 判断AliOssStore 对象是否实现Uploader接口
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint accessKey accessSecret has on empty...")
	}
	return nil
}

type AliOssStore struct {
	// 阿里云 OSS client, 私有变了， 不运行外部使用
	client *oss.Client
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		Endpoint:     os.Getenv("ALI_ANDPINT"),
		AccessKey:    os.Getenv("ALI_ACESS_KEY"),
		AccessSecret: os.Getenv("ALI_SECRET_KEY"),
	})

}
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	// 校验参数
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: c}, nil
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	// 获取ali oss bucket
	bucket, err := s.client.Bucket(bucketName)

	if err != nil {
		return err
	}

	// 上传文件到bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}

	// 获取文件的下载连接
	downloadUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载地址: %s\n", downloadUrl)
	fmt.Printf("地址有效期为24小时\n")
	return nil
}
