package main

import (
	"fmt"

	"github.com/zcalusic/sysinfo"
)

func main() {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	fmt.Println(si)
}
