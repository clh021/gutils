package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"path/filepath"
)

func testExt() {
	validFilename := GetValidScriptFile(".", "doSomeThing", "sh")
	log.Println("validFilename:", validFilename)
}

// 获取符合内定规则的有效脚本文件名
func GetValidScriptFile(dir, scriptName, scriptLastFix string) string {
	suffix := ".detect"
	lastFix := md5.Sum([]byte(scriptName + suffix + scriptLastFix))
	scriptFileName := fmt.Sprintf("%s.%s.%s", scriptName, hex.EncodeToString(lastFix[:])[:10], scriptLastFix)
	return filepath.Join(dir, "..", "scripts", scriptFileName)
}