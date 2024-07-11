package sshutil

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/povsister/scp"
)

func GetScpClient(ip string, port string, user string, password string, scpOption *scp.ClientOption) (*scp.Client, error) {
	scpConf := scp.NewSSHConfigFromPassword(user, password)
	client, err := scp.NewClient(ip+":"+port, scpConf, &scp.ClientOption{})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// RemoteRunScript 在远程主机上执行脚本，并返回执行结果
func RemoteRunScript(host, port, user, password, scriptPath string, args ...string) (string, string, error) {
	var stdoutBuf, stderrBuf bytes.Buffer
	scpClient, err := GetScpClient(host, port, user, password, &scp.ClientOption{})
	if err != nil {
		return "", "", fmt.Errorf("failed to create scpClient: %s", err)
	}
	remotePath := "/tmp/script.sh"
	err = scpClient.CopyFileToRemote(scriptPath, remotePath, &scp.FileTransferOption{Perm: os.FileMode(0755), Timeout: 60 * time.Second})
	if err != nil {
		return "", "", fmt.Errorf("failed to copy file: %s", err)
	}

	// client
	client, err := GetSshClient(host, port, user, password)
	if err != nil {
		return "", "", err
	}
	defer client.Close()

	// session
	session, err := client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session: %s", err)
	}
	defer session.Close()

	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	err = session.Run(fmt.Sprintf("%s %s", remotePath, strings.Join(args, " ")))
	if err != nil {
		return stdoutBuf.String(), stderrBuf.String(), fmt.Errorf("failed to execute command: %s", err)
	}

	return stdoutBuf.String(), stderrBuf.String(), nil
}

// 通过检查本地备用包是否存在，而增加备用包兼容性
// 拷贝本地文件到远程目录并授予非json文件为可执行权限
func CopyToRemote(host, port, user, password, localPath, remotePath string) error {
	scpClient, err := GetScpClient(host, port, user, password, &scp.ClientOption{})
	if err != nil {
		return fmt.Errorf("failed to create scpClient: %s", err)
	}

	localPathInfo, err := os.Stat(localPath)
	if err != nil {
		return err
	}

	if localPathInfo.IsDir() {
		return scpClient.CopyDirToRemote(localPath, remotePath, &scp.DirTransferOption{Timeout: 60 * time.Second}) //PreserveProp: true, protocol error: mode not delimited
	} else {
		return scpClient.CopyFileToRemote(localPath, remotePath, &scp.FileTransferOption{Perm: os.FileMode(0755), Timeout: 60 * time.Second})
	}
}
