package sshutil

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"

	"github.com/clh021/gutils/ufunc/envutil"
	"golang.org/x/crypto/ssh"
)

func GetSshClient(ip string, port string, user string, password string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
	}
	client, err := ssh.Dial("tcp", net.JoinHostPort(ip, port), config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// RemoteRun 在远程主机上执行脚本，并返回执行结果
func RemoteRun(ip string, port string, user string, password string, cmd string) (string, string, error) {
	client, err := NewSshClient(ip, port, user, password)
	if err != nil {
		return "", "", err
	}
	defer client.Close()
	return client.Run(cmd)
}

func RunScript(scriptFile string, args ...string) (string, string, error) {
	if !envutil.IsFileExist(scriptFile) {
		return "", "", fmt.Errorf("error: Can`t find script file %s", scriptFile)
	}
	cmd := exec.Command(scriptFile, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
