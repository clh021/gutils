# web

静态文件打包

## 自动生成静态文件，避免编译错误

```bash
go generate ./...
# 运行 generate.go 会自动生成静态文件
```

## 自动路由静态资源，支持 Vue 的 Hash路由

```bash
go run main.go
```