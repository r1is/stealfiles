package main

import (
	"encoding/json"
	"fmt"
	"upload_oss/common"
)

func main() {

	var argsInfo common.Args
	common.Flag(&argsInfo)
	// 检查argsInfo的值是否为空
	if argsInfo.Sm4key == "" || argsInfo.FileName == "" {
		// 退出程序
		return
	}
	// 打印参数
	// fmt.Println("key:", argsInfo.Sm4key)
	// fmt.Println("fileName:", argsInfo.FileName)

	var ossCfg common.OssCfg
	encText := "56821561bc48f67c52afac675347899fe58e0ce6369ab3b82362f3f794f068ee3aa08a58897f09626d0edbd183f72f7499a1329a77534eb9811363999e7a4fbba09ab2a48df8981ec200dd2974cd55d1b8f14b72b42eea15c7952433f86109b1ecfec43ddbd2b2cf9a477628f7543c1b2723685b2fde3baa33e6b6c636a58637e8fc9f6821b4a5c56b1636de107637f8375a6a0746dd19c300566bc6f2e40a6d168b6dcc4d85a849c975c903a6dc7c625577e939e51aa19a73cee9686b3ac50c1fac3d0147828c7b729c0bf2a0b17e343ab06819a655fb74d487520f8fdb2241aba241"
	// key := "1234567890abcdef"
	b := common.Sm4_d(argsInfo.Sm4key, encText)
	fmt.Println(b)
	// 将b赋值给jsonData
	err := json.Unmarshal([]byte(b), &ossCfg)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	common.Oss(ossCfg, argsInfo.FileName)

}
