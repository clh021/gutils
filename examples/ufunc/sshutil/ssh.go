package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/sshutil"
)

func testSsh() {
	stdOut, err := sshutil.RemoteRun("127.0.0.1", "8022", "root", "root", "ip a")
	if err != nil {
		log.Println(err)
	}
	log.Println("RemoteRun stdOut: ", stdOut)

	log.Println("----------------------")

	stdOut, stdErr, err := sshutil.RunScript("./test.sh", "arg1", "arg2")
	if err != nil {
		log.Println("RunScript err: ", err)
	}
	log.Println("RunScript stdOut: ", stdOut)
	log.Println("RunScript stdErr: ", stdErr)
}