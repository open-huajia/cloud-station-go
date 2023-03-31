package store

// 定义如何上传文件到bucket
// 抽象为接口, 不关心我们使用的是哪个OSS bucket
type Uploader interface {
	Upload(bucketName string, objectKey string, fileName string) error
}
