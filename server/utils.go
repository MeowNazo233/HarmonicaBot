package server

import (
	"os"
)

// 判断文件夹是否存在 若不存在则新建
func CheckPath(path string) error {
	exist, err := pathExists(path)
	if !exist && err == nil {
		return (os.MkdirAll(path, os.ModePerm))
	}
	return err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
