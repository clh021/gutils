package web

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type AssetsServer struct {
	embedFS_web      embed.FS
	embedFS_rootpath string // TODO: support disk local has file to use
}

func NewAssetsServer(embedFS_web embed.FS, rootpath string) *AssetsServer {
	return &AssetsServer{
		embedFS_web:      embedFS_web,
		embedFS_rootpath: rootpath, // "assets/dist",
	}
}

func (s AssetsServer) getContentType(extension string) string {
	contentTypes := map[string]string{
		".css":  "text/css",
		".js":   "application/javascript",
		".html": "text/html; charset=utf-8",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		// 添加更多类型...
	}
	return contentTypes[extension] // 如果未找到，将返回零值，即默认""
}

func (s AssetsServer) Dump() {
	err := fs.WalkDir(s.embedFS_web, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (s AssetsServer) serveIndexPath(w http.ResponseWriter, embedFS_dirname string) {
	indexHtml, err := s.embedFS_web.ReadFile(filepath.Join(embedFS_dirname, "/index.html"))
	if err != nil {
		log.Printf("Error reading index.html: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write(indexHtml)
}

// http.HandleFunc("/", assets.RouteWeb) // 统一处理器处理所有请求
// log.Printf("Server started at :%d", config.Port)
// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
func (s AssetsServer) RouteWeb(w http.ResponseWriter, r *http.Request) {
	relativePath := path.Clean(r.URL.Path[len("/"):])
	if relativePath == "." {
		s.serveIndexPath(w, s.embedFS_rootpath)
		return
	}
	if strings.Contains(relativePath, "..") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	log.Println("relativePath file:", relativePath)
	filePath := filepath.Join(s.embedFS_rootpath, relativePath)

	log.Println("Requested file:", filePath)
	reader, err := s.embedFS_web.Open(filePath)
	if err != nil {
		// 静态文件不存在时，如果是 assets 开头的请求路径则直接返回 404 否则返回 index.html
		if strings.HasPrefix(r.URL.Path, "assets") {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			s.serveIndexPath(w, s.embedFS_rootpath)
		}
		return
	}
	defer reader.Close()

	// 获取文件扩展名并设置Content-Type
	ext := filepath.Ext(filePath)
	contentType := s.getContentType(ext)
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	} else {
		// 如果未知类型，默认处理
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	_, err = io.Copy(w, reader)
	if err != nil {
		log.Printf("Error serving file: %v", err)
		http.Error(w, "Error serving file", http.StatusInternalServerError)
	}
}

func GenerateFile(filePath, content string) error {
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
