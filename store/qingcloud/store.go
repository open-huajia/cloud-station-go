package qingcloud

func NewQingCloudOssStore() *QingCloudOssStore {
	return &QingCloudOssStore{}
}

type QingCloudOssStore struct {
}

func (q *QingCloudOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
