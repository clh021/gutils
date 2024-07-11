package config

import (
	"fmt"
	"log"

	"github.com/clh021/gutils/ufunc/runtimehelpers"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Load 初始化并加载配置
func Load(Config interface{}) error {
	viper.SetConfigName("config") // 配置文件名（无扩展名）
	viper.SetConfigType("yaml")   // 假设配置文件是 YAML 格式
	viper.AddConfigPath(".")      // 查找配置文件的路径
	viper.AddConfigPath(runtimehelpers.GetProgramPath())

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果配置文件不存在，直接退出程序
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("config file not found")
		} else {
			return fmt.Errorf("fatal error loading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(Config); err != nil {
		return err
	}
	viper.OnConfigChange(func(e fsnotify.Event) {

		log.Printf("Config file changed: %s Op:%s\n", e.Name, e.Op)
		if err := viper.Unmarshal(Config); err != nil {
			log.Printf("Error unmarshaling config on change: %v\n", err)
		}
	})

	log.Println("Using config file:", viper.ConfigFileUsed())
	log.Println("setting init success !")
	return nil
}
