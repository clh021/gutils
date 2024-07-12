package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/sshutil"
)

func testSshClient() {
	client, err := sshutil.NewSshClient("127.0.0.1", "8022", "root", "root")
	if err != nil {
		log.Println(err)
	}
	defer client.Close()
	stdOut, stdErr, err := client.Run("ip a")
	if err != nil {
		log.Println(err)
	}
	log.Println("SshClient Run stdOut: ", stdOut)
	log.Println("SshClient Run stdErr: ", stdErr)
}