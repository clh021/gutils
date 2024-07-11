package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/envutil"
)

func testFile() {
	if isExist := envutil.IsFileExist("./README.md"); isExist {
		log.Println("IsFileExist:", isExist)
	} else {
		log.Println("IsFileExist:", isExist)
	}

	findFiles, err := envutil.FindFilesWithPrefixAndSuffix(".", "p", ".go")
	if err != nil {
		log.Println("FindFilesWithPrefixAndSuffix err:", err)
	} else {
		log.Println("FindFilesWithPrefixAndSuffix:", findFiles)
	}

	if envutil.IsCmdExists("ls") {
		log.Println("IsCmdExists:", "ls exist")
	} else {
		log.Println("IsCmdExists:", "ls not exist")
	}

	if envutil.IsCmdExists("arp") {
		log.Println("IsCmdExists:", "arp exist")
	} else {
		log.Println("IsCmdExists:", "arp not exist")
	}

	file := envutil.FindExistingFile([]string{"test.md", "./main.go", "./README.md"})
	log.Println("FindExistingFile:", file)

}
