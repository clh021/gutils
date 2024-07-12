package sshutil

import (
	"bytes"
	"fmt"
	"strings"

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

// IsFile 检查远程服务器上的路径是否为文件
func (c *SshClient) IsFile(path string) (bool, error) {
	_, _, err := c.Run(fmt.Sprintf("test -f %s", path))
	if err != nil {
		if exitError, ok := err.(*ssh.ExitError); ok {
			// Check if the command exited with a non-zero status
			if exitError.ExitStatus() == 1 {
				return false, nil // 文件不存在或不是文件
			}
		}
		return false, err // 其他类型的错误
	}
	return true, nil // 文件存在且是文件
}

// IsDir 检查远程服务器上的路径是否为目录
func (c *SshClient) IsDir(path string) (bool, error) {
	_, _, err := c.Run(fmt.Sprintf("test -d %s", path))
	if err != nil {
		if exitError, ok := err.(*ssh.ExitError); ok {
			if exitError.ExitStatus() == 1 {
				return false, nil // 目录不存在或不是目录
			}
		}
		return false, err // 其他类型的错误
	}
	return true, nil // 目录存在且是目录
}

// IsRunning 检查远程服务器上是否有指定名称的进程正在运行
func (c *SshClient) IsRunning(processName string) (bool, error) {
	stdout, _, err := c.Run(fmt.Sprintf("pgrep %s", processName))
	if err != nil {
		if strings.Contains(err.Error(), "no process found") {
			return false, nil // 没有找到进程，不是错误
		}

		// 使用ps和grep组合来查找进程
		cmd := fmt.Sprintf("ps aux | grep '%s' | grep -v 'grep'", processName)
		stdout, _, err = c.Run(cmd)
		if err != nil {
			return false, err
		}
	}
	return len(strings.TrimSpace(stdout)) > 0, nil
}