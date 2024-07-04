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

var (
	Config = new(Conf)
)

func main() {
	var conf Conf
	err := config.Load(&conf)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	// 使用 spew.Dump 打印配置结构体，它提供了更详细的输出
	spew.Dump(conf)
	// 或者使用 fmt.Printf 格式化输出
	// fmt.Printf("Config: PrometheusAddress=%s, Port=%d\n", conf.PrometheusAddress, conf.Port)
}