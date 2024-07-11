package main

import (
	"fmt"
	"log"
	"time"

	"github.com/clh021/gutils/ufunc/sshutil"
)

func testScp() {
	var err error
	var stdOut string
	stdOut, stdErr, err := sshutil.RemoteRunScript("127.0.0.1", "8022", "root", "root", "./test.sh", "arg1", "arg2")
	if err != nil {
		log.Println("RemoteRunScript err: ", err)
	}
	log.Println("RemoteRunScript stdOut: ", stdOut)
	log.Println("RemoteRunScript stdErr: ", stdErr)

	log.Println("----------------------")

	err = sshutil.CopyToRemote("127.0.0.1", "8022", "root", "root", "./test.sh", fmt.Sprintf("%s%s", "/tmp/test.sh", time.Now().Format("20060102150405")))
	if err != nil {
		log.Println(err)
	}

	err = sshutil.CopyToRemote("127.0.0.1", "8022", "root", "root", "./testDir", fmt.Sprintf("%s%s", "/root/testCopyDir", time.Now().Format("20060102150405")))
	if err != nil {
		log.Println(err)
	}
	stdOut, err = sshutil.RemoteRun("127.0.0.1", "8022", "root", "root", "tree /tmp")
	if err != nil {
		log.Println(err)
	}
	log.Println("CopyToRemote test files stdOut: ", stdOut)
	log.Println("----------------------")
	stdOut, err = sshutil.RemoteRun("127.0.0.1", "8022", "root", "root", "tree /root")
	if err != nil {
		log.Println(err)
	}
	log.Println("CopyToRemote test dirs stdOut: ", stdOut)
}