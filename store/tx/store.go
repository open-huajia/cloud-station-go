package tx

func NewTxOssStore() *TxOssStore {
	return &TxOssStore{}
}

type TxOssStore struct {
}

func (t *TxOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
