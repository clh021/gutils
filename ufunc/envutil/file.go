package envutil

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsFileExist(p string) bool {
	info, err := os.Stat(p)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// 查找指定目录下指定前缀和后缀的文件
func FindFilesWithPrefixAndSuffix(dir, prefix, suffix string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，忽略目录
		if !info.Mode().IsRegular() {
			return nil
		}

		// 匹配文件名
		if len(prefix) > 0 || len(suffix) > 0 {
			if strings.HasPrefix(info.Name(), prefix) && strings.HasSuffix(info.Name(), suffix) {
				files = append(files, path)
			}
		} else if (len(prefix) > 0) {
			if strings.HasPrefix(info.Name(), prefix) {
				files = append(files, path)
			}
		} else if (len(suffix) > 0) {
			if strings.HasSuffix(info.Name(), suffix) {
				files = append(files, path)
			}
		} else {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func IsCmdExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// 判断一个文件路径数组中首先发现存在的文件
// 一般用来从多个可能允许存在的路径中，按照先后顺序(优先级)依次查找发现的存在。
func FindExistingFile(paths []string) string {
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}