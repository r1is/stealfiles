package common

import "os"

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
