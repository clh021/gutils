package sshutil

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/povsister/scp"
)

type ScpClient struct {
	client    *scp.Client
	SshClient *SshClient
	_ip       string
	_port     string
	_user     string
	_password string
}

// NewScpClient 创建一个scp客户端
// 适用于 一次连接 希望执行多条命令的情况
// 本客户端内嵌套了ssh客户端，所以使用本客户端时，可以复用ssh客户端，不用额外创建 ssh 客户端。
func NewScpClient(ip string, port string, user string, password string) (*ScpClient, error) {
	client, err := GetScpClient(ip, port, user, password, &scp.ClientOption{})
	if err != nil {
		return nil, err
	}
	return &ScpClient{client: client, _ip: ip, _port: port, _user: user, _password: password}, nil
}

func (c *ScpClient) Close() (err error) {
	if c.SshClient == nil {
		err = nil
	} else {
		err = c.SshClient.Close()
	}
	return
}

func (c *ScpClient) InitSshClient() (err error) {
	if c.SshClient == nil {
		c.SshClient, err = NewSshClient(c._ip, c._port, c._user, c._password)
	}
	return err
}

func (c *ScpClient) Run(cmd string) (string, string, error) {
	return c.SshClient.Run(cmd)
}

func (c *ScpClient) RunScript(localScriptPath string, args ...string) (string, string, error) {
	remotePath := "/tmp/script.sh"
	err := c.client.CopyFileToRemote(localScriptPath, remotePath, &scp.FileTransferOption{Perm: os.FileMode(0755), Timeout: 60 * time.Second})
	if err != nil {
		return "", "", fmt.Errorf("failed to copy file: %s", err)
	}
	err = c.InitSshClient()
	if err != nil {
		return "", "", err
	}
	return c.SshClient.Run(fmt.Sprintf("%s %s", remotePath, strings.Join(args, " ")))
}

func (c *ScpClient) Copy(localPath string, remotePath string) error {
	localPathInfo, err := os.Stat(localPath)
	if err != nil {
		return err
	}

	if localPathInfo.IsDir() {
		return c.client.CopyDirToRemote(localPath, remotePath, &scp.DirTransferOption{Timeout: 60 * time.Second}) //PreserveProp: true, protocol error: mode not delimited
	} else {
		return c.client.CopyFileToRemote(localPath, remotePath, &scp.FileTransferOption{Perm: os.FileMode(0755), Timeout: 60 * time.Second})
	}
}
