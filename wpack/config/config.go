package config

import (
	"fmt"
	"log"

	"github.com/clh021/gutils/ufunc/runtimehelpers"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ConfigLoadOptions 定义配置加载的选项结构体，包含默认值
type LoadOpts struct {
	Paths   []string // 默认配置路径
	CName   string   `default:"config"` // 默认配置文件名
	CType   string   `default:"yaml"`   // 默认配置文件类型
}

// Load 初始化并加载配置
// params:
// Config 配置结构体
// paths 配置文件路径 eg:[".","config","conf"]
// cname 配置文件名 eg:"config"
// ctype 配置文件类型 eg:"yaml"
func Load(Config interface{}, opts *LoadOpts) (string, error) {
		// 设置默认值（如果用户没有覆盖的话）
	if len(opts.Paths) == 0 {
		opts.Paths = []string{"./", "config", "conf"}
	}
	if opts.CName == "" {
		opts.CName = "config"
	}
	if opts.CType == "" {
		opts.CType = "yaml"
	}

	viper.SetConfigName(opts.CName) // 配置文件名（无扩展名）
	viper.SetConfigType(opts.CType) // 假设配置文件是 YAML 格式
	for _, p := range opts.Paths {
		viper.AddConfigPath(p) // 查找配置文件的路径
	}
	viper.AddConfigPath(runtimehelpers.GetProgramPath())

	confPath := ""

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果配置文件不存在，直接退出程序
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return confPath, fmt.Errorf("config file not found")
		} else {
			return confPath, fmt.Errorf("fatal error loading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(Config); err != nil {
		return confPath, err
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s Op:%s\n", e.Name, e.Op)
		if err := viper.Unmarshal(Config); err != nil {
			log.Printf("Error unmarshaling config on change: %v\n", err)
		}
	})

	confPath = viper.ConfigFileUsed()
	return confPath, nil
}
