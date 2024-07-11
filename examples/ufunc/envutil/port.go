package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/envutil"
)

func testPort() {
	port, err := envutil.GetFreePort()
	if err != nil {
		log.Println("GetFreePort err:", err)
	} else {
		log.Println("GetFreePort :", port)
	}

}