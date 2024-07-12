package sshutil

import (
	"fmt"

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
	client, err := NewScpClient(host, port, user, password)
	if err != nil {
		return "", "", fmt.Errorf("failed to create scpClient: %s", err)
	}
	defer client.Close()
	return client.RunScript(scriptPath, args...)
}

// 通过检查本地备用包是否存在，而增加备用包兼容性
// 拷贝本地文件到远程目录并授予非json文件为可执行权限
func CopyToRemote(host, port, user, password, localPath, remotePath string) error {
	client, err := NewScpClient(host, port, user, password)
	if err != nil {
		return fmt.Errorf("failed to create scpClient: %s", err)
	}
	defer client.Close()
	return client.Copy(localPath, remotePath)
}
