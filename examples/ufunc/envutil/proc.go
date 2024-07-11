package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/envutil"
)

func testProc() {
	owner := envutil.GetProcessOwner()
	log.Println("GetProcessOwner: ", owner)
}