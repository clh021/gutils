package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/clh021/gutils/wpack/web"
)

// eg1: go:generate go run assets/generate_helper.go
// eg2:
//go:generate mkdir -p assets/dist
//go:generate cp -n assets_gen_tpl/index.html assets/dist/index.html
//go:generate cp -rn assets_gen_tpl/assets assets/dist/
//go:generate touch assets/dist/.gitkeep

//go:embed assets/dist/*
var embedFS_web embed.FS

func main() {
	// wpack.RouteWeb(embedFS_web)
	port := 8080
	assets := web.NewAssetsServer(embedFS_web, "assets/dist")
	assets.Dump()
	http.HandleFunc("/", assets.RouteWeb)
	log.Printf("Server started at :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
