package main

import (
	"fmt"

	"github.com/clh021/gutils/ufunc/exec"
)

func main() {
	p := exec.GetProgramPath()
	fmt.Println("GetProgramPath:", p)
}
