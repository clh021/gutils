package main

import (
	"fmt"

	"github.com/clh021/gutils/wpack/devenvinfo"
)


func main() {
	devinfo := devenvinfo.GetDevelopments()
	fmt.Println("devinfo:", devinfo)
}