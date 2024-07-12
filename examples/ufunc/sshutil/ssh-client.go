package main

import (
	"log"

	"github.com/clh021/gutils/ufunc/sshutil"
)

func testSshClient() {
	client, err := sshutil.NewSshClient("127.0.0.1", "8022", "root", "root")
	if err != nil {
		log.Println("sshutil.NewSshClient:", err)
	}
	defer client.Close()
	stdOut, stdErr, err := client.Run("ip a")
	if err != nil {
		log.Println("SshClient Run err: ", err)
	}
	log.Println("SshClient Run stdOut: ", stdOut)
	log.Println("SshClient Run stdErr: ", stdErr)

	stdOut, stdErr, err = client.Run("mkdir -p /root/testDir; touch /root/testDir/testFile")
	if err != nil {
		log.Println("SshClient Run err: ", err)
	}
	log.Println("SshClient Run stdOut: ", stdOut)
	log.Println("SshClient Run stdErr: ", stdErr)

	if isFile, err := client.IsFile("/root/testDir/testFile"); err != nil {
		log.Println("isFile: /root/testDir/testFile : false")
	} else {
		log.Println("isFile: /root/testDir/testFile :", isFile)
	}
	if isDir, err := client.IsDir("/root/testDir"); err != nil {
		log.Println("IsDir err: ", err)
	} else {
		log.Println("isDir: /root/testDir :", isDir)
	}
	if isFile, err := client.IsFile("/root/testDir/testFile7777"); err != nil {
		log.Println("IsFile err: ", err)
	} else {
		log.Println("isFile: /root/testDir/testFile :", isFile)
	}
	if isDir, err := client.IsDir("/root/testDir7777"); err != nil {
		log.Println("IsDir err: ", err)
	} else {
		log.Println("isDir: /root/testDir :", isDir)
	}
	if isRunning, err := client.IsRunning("sshd"); err != nil {
		log.Println("isRunning err: ", err)
	} else {
		log.Println("isRunning : sshd: ", isRunning)
	}
	if cmdPath, err := client.FindCmdPath("sshd"); err != nil {
		log.Println("FindCmdPath err: ", err)
	} else {
		log.Println("FindCmdPath : sshd: ", cmdPath)
	}
	if cmdPath, err := client.FindCmdPath("ash"); err != nil {
		log.Println("FindCmdPath err: ", err)
	} else {
		log.Println("FindCmdPath : ash: ", cmdPath)
	}
	if cmdPath, err := client.FindCmdPath("bash"); err != nil {
		log.Println("FindCmdPath err: ", err)
	} else {
		log.Println("FindCmdPath : bash: ", cmdPath)
	}
}
