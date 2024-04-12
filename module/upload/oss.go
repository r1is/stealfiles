package upload

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
	"upload_oss/common"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// 只能上传文件
func UploadFile(ossCfg common.OssCfg, filePath string) (string, error) {
	u, _ := url.Parse(ossCfg.BatchURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:     ossCfg.TmpSecretID,
			SecretKey:    ossCfg.TmpSecretKey,
			SessionToken: ossCfg.SessionToken,
		},
	})
	fleAbsPath, _ := filepath.Abs(filePath)
	fileName := filepath.Base(fleAbsPath)
	fmt.Println("fileName:", fileName)
	fmt.Println("fleAbsPath:", fleAbsPath)

	if is, err := IsFile(fleAbsPath); err != nil {
		return "", err
	} else {
		if is {
			// 从命令行获取参数
			key := "file/" + fileName //对象键（Key）是对象在存储桶中的唯一标识

			opt := &cos.MultiUploadOptions{ThreadPoolSize: 8, CheckPoint: true}
			resp, _, err := client.Object.Upload(
				context.Background(), key, fleAbsPath, opt,
			)
			if err != nil {
				return "", err
			}
			return resp.Key, nil
		} else {
			return "", errors.New("this is a folder not a file")
		}
	}
}

// 通过 tag 的方式，用户可以将请求参数或者请求头部放进签名中。
type URLToken struct {
	SessionToken string `url:"x-cos-security-token,omitempty" header:"-"`
}

func GetPresignedURL(ossCfg common.OssCfg, key string) {
	tak := ossCfg.TmpSecretID
	tsk := ossCfg.TmpSecretKey
	token := &URLToken{
		SessionToken: ossCfg.SessionToken,
	}

	u, _ := url.Parse(ossCfg.BatchURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})
	ctx := context.Background()

	// 方法2 通过 tag 设置 x-cos-security-token
	// 获取预签名
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodGet, key, tak, tsk, 10*time.Minute, token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("DownloadRUL: ", presignedURL.String())

}

// 在Linux下判断一个文件是否是文件夹
func IsFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if fileInfo.IsDir() {
		return false, nil
	} else {
		return true, nil
	}
}

// func _zip(dst, src string) (err error) {
// 	// 创建准备写入的文件
// 	fw, err := os.Create(dst)
// 	defer fw.Close()
// 	if err != nil {
// 		return err
// 	}

// 	// 通过 fw 来创建 zip.Write
// 	zw := zip.NewWriter(fw)
// 	defer func() {
// 		// 检测一下是否成功关闭
// 		if err := zw.Close(); err != nil {
// 			log.Fatalln(err)
// 		}
// 	}()

// 	// 下面来将文件写入 zw ，因为有可能会有很多个目录及文件，所以递归处理
// 	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
// 		if errBack != nil {
// 			return errBack
// 		}

// 		// 通过文件信息，创建 zip 的文件信息
// 		fh, err := zip.FileInfoHeader(fi)
// 		if err != nil {
// 			return
// 		}

// 		// 替换文件信息中的文件名
// 		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

// 		// 这步开始没有加，会发现解压的时候说它不是个目录
// 		if fi.IsDir() {
// 			fh.Name += "/"
// 		}

// 		// 写入文件信息，并返回一个 Write 结构
// 		w, err := zw.CreateHeader(fh)
// 		if err != nil {
// 			return
// 		}

// 		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
// 		// 如目录，也没有数据需要写
// 		if !fh.Mode().IsRegular() {
// 			return nil
// 		}

// 		// 打开要压缩的文件
// 		fr, err := os.Open(path)
// 		defer fr.Close()
// 		if err != nil {
// 			return
// 		}

// 		// 将打开的文件 Copy 到 w
// 		_, err1 := io.Copy(w, fr)
// 		if err1 != nil {
// 			return
// 		}

// 		return nil
// 	})
// }
