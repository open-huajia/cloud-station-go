package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/open-huajia/cloud-station-go/store"
	"os"
)

var (
	// 对象是否实现了接口的约束
	// a string = "abc"
	// _ store.Uploader 我不需要这个变量的值, 我只是做变量类型的判断
	// &AliOssStore{} 这个对象 必须满足 store.Uploader
	// _ store.Uploader = &AliOssStore{} 声明了一个空对象, 只是需要一个地址
	// nil 空指针, nil有没哟类型: 有类型
	// a *AliOssStore = nil   nil是一个AliOssStore 的指针
	// 如何把nil 转化成一个 指定类型的变量
	//    a int = 16
	//    b int64 = int64(a)
	//    (int64类型)(值)
	//	  (*AliOssStore)(nil)
	_ store.Uploader = (*AliOssStore)(nil)
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
	client   *oss.Client
	listener oss.ProgressListener
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
	return &AliOssStore{client: c, listener: NewDefaultProgressListener()}, nil
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	// 获取ali oss bucket
	bucket, err := s.client.Bucket(bucketName)

	if err != nil {
		return err
	}

	// 上传文件到bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName, oss.Progress(s.listener)); err != nil {
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
