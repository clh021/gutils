package main

import (
	"fmt"

	"github.com/clh021/gutils/wpack/config"
	"github.com/davecgh/go-spew/spew"
)


type Conf struct {
	PrometheusAddress string `mapstructure:"prometheus_address"`
	Port              int    `mapstructure:"port"`
}

func main() {
	var conf Conf
	confPath, err := config.Load(&conf, &config.LoadOpts{})
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	} else {
		fmt.Printf("Loaded config from %s\n", confPath)
	}

	// 使用 spew.Dump 打印配置结构体，它提供了更详细的输出
	spew.Dump(conf)
}