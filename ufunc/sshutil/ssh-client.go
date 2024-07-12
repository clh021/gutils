package sshutil

import (
	"bytes"

	"golang.org/x/crypto/ssh"
)

type SshClient struct {
	client  *ssh.Client
}

// NewSshClient 创建一个ssh客户端
// 适用于 一次连接 希望执行多条命令的情况
func NewSshClient(ip string, port string, user string, password string) (*SshClient, error) {
	client, err := GetSshClient(ip, port, user, password)
	if err != nil {
		return nil, err
	}
	return &SshClient{client: client}, nil
}

func (c *SshClient) Close() error {
	return c.client.Close()
}

func (c *SshClient) Run(cmd string) (string, string, error) {
	// 一个 session 只对一个 Run
	session, err := c.client.NewSession()
	if err != nil {
		return "", "", err
	}
	defer session.Close()
	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	// Run runs cmd on the remote host. Typically, the remote server passes cmd to the shell for interpretation. A Session only accepts one call to Run, Start, Shell, Output, or CombinedOutput.
	err = session.Run(cmd)
	return stdoutBuf.String(), stderrBuf.String(), err
}
