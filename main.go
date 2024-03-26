package main

import (
	"fmt"
	"upload_oss/common"
	"upload_oss/utils"
)

func main() {
	var argsInfo common.Args
	common.Flag(&argsInfo)
	// 检查argsInfo的值是否为空
	if argsInfo.Passcode == "" || argsInfo.FileName == "" {
		// 退出程序
		return
	}

	targetURL := utils.GetServerURL()
	ossCfg, err := utils.GetOssCfg(targetURL, argsInfo)
	if err != nil {
		return
	}
	key, err := common.UploadFile(ossCfg, argsInfo)
	if err != nil {
		fmt.Println("UploadFile failed:", err)
		return
	}
	//获取预签名URL
	common.GetPresignedURL(ossCfg, key)
}
