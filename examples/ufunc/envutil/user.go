package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/envutil"
)

func testUser() {
	if isSuper, err := envutil.CheckSuperUser(); err != nil {
		log.Println("CheckSuperUser err:", err)
	} else {
		log.Println("CheckSuperUser:", isSuper)
	}
}