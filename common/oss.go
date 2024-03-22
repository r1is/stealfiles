package common

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func Oss(ossCfg OssCfg, fileName string) {
	var SECRETID = ossCfg.Ak
	var SECRETKEY = ossCfg.Sk

	u, _ := url.Parse(ossCfg.BatchURL)

	su, _ := url.Parse(ossCfg.BucketURL)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SECRETID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: SECRETKEY, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	name := "test/"
	// 从命令行获取参数
	name += fileName
	// 2.通过本地文件上传对象
	_, err := c.Object.PutFromFile(context.Background(), name, fileName, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("上传成功,download url：", ossCfg.BatchURL+"/"+name)
}
