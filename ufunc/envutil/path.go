package envutil

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)


func GetProgramPath() string {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exReal)
}

func IsDirExist(p string) bool {
	info, err := os.Stat(p)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// 计算目录的哈希值
func GetDirHash(dirPath string) (string, error) {
    // 获取目录下的所有文件名
    files, err := os.ReadDir(dirPath)
    if err != nil {
        return "", err
    }

    // 对文件名进行排序
    fileNames := make([]string, 0, len(files))
    for _, file := range files {
        if !file.IsDir() {
            fileNames = append(fileNames, file.Name())
        }
    }
    sort.Strings(fileNames)

    // 计算所有文件内容的哈希值
    h := md5.New()
    for _, fileName := range fileNames {
        filePath := filepath.Join(dirPath, fileName)
        fileBytes, err := os.ReadFile(filePath)
        if err != nil {
            return "", err
        }
        h.Write(fileBytes)
    }
    hashBytes := h.Sum(nil)

    // 将哈希值转换为字符串
    hashStr := fmt.Sprintf("%x", hashBytes)
    return hashStr, nil
}

// 获取目录的修改时间
func GetDirModTime(dirPath string) (string, error) {
    // 获取目录信息
    dirInfo, err := os.Stat(dirPath)
    if err != nil {
        return "", err
    }

    // 获取目录的修改时间
    return dirInfo.ModTime().Format("20060102.150405"), nil
}