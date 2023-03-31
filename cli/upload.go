package cli

import (
	"github.com/open-huajia/cloud-station-go/store"
	"github.com/open-huajia/cloud-station-go/store/aliyun"
	"github.com/open-huajia/cloud-station-go/store/aws"
	"github.com/open-huajia/cloud-station-go/store/qingcloud"
	"github.com/open-huajia/cloud-station-go/store/tx"
	"github.com/spf13/cobra"
	"os"
)

var (
	default_ak = os.Getenv("ALI_ACESS_KEY")
	default_sk = os.Getenv("ALI_SECRET_KEY")
)

var (
	ossProvier   string
	ossEndpoint  string
	accessKey    string
	accessSecret string
	bucketName   string
	uploadFile   string
)

func setAliDefault(opts *aliyun.Options) {
	if opts.AccessKey == "" {
		opts.AccessKey = default_ak
	}
	if opts.AccessSecret == "" {
		opts.AccessSecret = default_sk
	}
}

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "upload file",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvier {
		case "aliyun":
			aliOpts := &aliyun.Options{
				Endpoint:     ossEndpoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			}
			setAliDefault(aliOpts)
			uploader, err = aliyun.NewAliOssStore(aliOpts)

		case "tx":
			uploader = tx.NewTxOssStore()
		case "aws":
			uploader = aws.NewAwsOssStore()
		case "qingcload":
			uploader = qingcloud.NewQingCloudOssStore()
		}
		if err != nil {
			return err
		}
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvier, "provider", "p", "aliyun", "oss storage provier [aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provier endpoint")
	f.StringVarP(&bucketName, "bucket_name", "b", "devgo-station", "oss storage provier bucket name")
	f.StringVarP(&accessKey, "access_key", "k", "", "oss storage provier ak")
	f.StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provier sk")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}
