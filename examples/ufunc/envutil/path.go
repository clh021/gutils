package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/envutil"
)

func testPath() {
	p := envutil.GetProgramPath()
	log.Println("GetProgramPath:", p)

	if isExist := envutil.IsDirExist("."); isExist {
		log.Println("IsDirExist:", isExist)
	} else {
		log.Println("IsDirExist:", isExist)
	}

	if hash, err := envutil.GetDirHash("."); err != nil {
		log.Println("GetDirHash err:", err)
	} else {
		log.Println("GetDirHash:", hash)
	}

	if hash, err := envutil.GetDirModTime("."); err != nil {
		log.Println("GetDirModTime err:", err)
	} else {
		log.Println("GetDirModTime:", hash)
	}
}