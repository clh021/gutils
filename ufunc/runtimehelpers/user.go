package runtimehelpers

import (
	"fmt"
	"os/user"
)

// CheckSuperUser 权限检查函数，用于验证当前用户是否为超级用户。
// 如果不是超级用户，返回一个具体的错误信息。
// 如果遇到其他错误，也返回相应的错误描述。
// 该函数不会导致程序退出或产生标准输出，仅返回错误。
func CheckSuperUser() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, fmt.Errorf("获取当前用户信息失败: %w", err)
	}

	return currentUser.Uid == "0", nil
}
