package main

import (
	"fmt"
	"log"
	"time"

	"github.com/clh021/gutils/ufunc/sshutil"
)

func testScpClient() {
	nowTime := time.Now().Format("20060102150405")
	client, err := sshutil.NewScpClient("127.0.0.1", "8022", "root", "root")
	if err != nil {
		log.Println("ScpClient RemoteRunScript err: ", err)
	}
	defer client.Close()

	stdOut, stdErr, err := client.RunScript("test.sh", "arg1", "arg2")
	if err != nil {
		log.Println("ScpClient RunScript err: ", err)
	}
	log.Println("ScpClient RunScript stdOut: ", stdOut)
	log.Println("ScpClient RunScript stdErr: ", stdErr)
	log.Println("----------------------")

	err = client.Copy("./test.sh", fmt.Sprintf("%s%s", "/tmp/test.sh", nowTime))
	if err != nil {
		log.Println("ScpClient Copy file err: ", err)
	} else {
		log.Println("ScpClient Copy file ok: ", err)
	}

	err = client.Copy("./testDir", fmt.Sprintf("%s%s", "/root/testCopyDir", nowTime))
	if err != nil {
		log.Println("ScpClient Copy dir err:", err)
	} else {
		log.Println("ScpClient Copy dir ok:", err)
	}
	stdOut, stdErr, err = client.Run("tree /tmp")
	if err != nil {
		log.Println("ScpClient.SshClient.Run err:", err)
	}
	log.Println("ScpClient Copy test files stdOut: ", stdOut)
	log.Println("ScpClient Copy test files stdErr: ", stdErr)
	log.Println("----------------------")
	stdOut, stdErr, err = client.Run("tree /root")
	if err != nil {
		log.Println("ScpClient.SshClient.Run err:", err)
	}
	log.Println("ScpClient Copy test dirs stdOut: ", stdOut)
	log.Println("ScpClient Copy test dirs stdErr: ", stdErr)
}
