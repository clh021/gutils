//go:build ignore
// +build ignore

//go:generate echo $GOFILE
//go:generate go run $GOFILE

package main

import (
	"os"
	"path/filepath"
)

func main() {
	// 生成index.html
	htmlContent := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/index.css">
    <title>Index</title>
</head>
<body>
    这是 Index 页面。
</body>
</html>
`
	generateFile("./dist/index.html", htmlContent)

	// 生成assets/index.css
	cssContent := `
body {
    font-family: Arial, sans-serif;
    color: #333;
    margin: 20px;
}
`
	generateFile("./dist/assets/index.css", cssContent)
}

func generateFile(filePath, content string) error {
	// 检查目录是否存在，如果不存在则创建
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// 检查文件是否存在，如果不存在则创建并写入内容
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}