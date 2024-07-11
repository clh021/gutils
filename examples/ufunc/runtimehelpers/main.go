package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/runtimehelpers"
)

func main() {
	p := runtimehelpers.GetProgramPath()
	log.Println("GetProgramPath:", p)

	port, err := runtimehelpers.GetFreePort()
	if err != nil {
		log.Println("err:", err)
	} else {
		log.Println("port:", port)
	}

	if isSuper, err := runtimehelpers.CheckSuperUser(); err != nil {
		log.Println("err:", err)
	} else {
		log.Println("CheckSuperUser:", isSuper)
	}
}
