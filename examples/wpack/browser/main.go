package main

import (
	"log"

	"github.com/clh021/gutils/wpack/browser"
)


func main() {
	bs := browser.GetBrowsers()
	log.Println(*bs)
}