package aliyun_test

import (
	"github.com/open-huajia/cloud-station-go/store"
	"github.com/open-huajia/cloud-station-go/store/aliyun"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	uploader store.Uploader
)

var (
	bucketName = os.Getenv("ALI_BUCKET_NAME")
)

func TestUpload(t *testing.T) {
	// 使用assert编写测试用例断言语句
	// 通过New获取一个断言实例
	should := assert.New(t)
	err := uploader.Upload(bucketName, "test.txt", "store_test.go")

	if should.NoError(err) {
		t.Log("upload ok")
	}
}

func TestUploadError(t *testing.T) {
	should := assert.New(t)

	err := uploader.Upload(bucketName, "test.txt", "store_testxxx.go")
	should.Error(err, "open store_testxxx.go: The system cannot find the file specified.")
}

func init() {
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali

}
