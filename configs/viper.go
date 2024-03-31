package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"wechatbot/log"
)

func InitConfig() {

	//加载配置文件位置
	viper.SetConfigFile("configs/config.yaml")
	//监听配置文件
	viper.WatchConfig()
	//监听是否更改配置文件
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Logger.Info("config is change")
		err := viper.Unmarshal(&Cfg)
		if err != nil {
			log.Logger.Error("config unmarshal fail", zap.Error(err))
			return
		}
	})
	// 读取配置文件内容
	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.Error("read config fail", zap.Error(err))
		return
	}

	//将配置文件内容写入到Conf结构体
	if err = viper.Unmarshal(&Cfg); err != nil {
		log.Logger.Error("config unmarshal fail", zap.Error(err))
		return
	}

	return
}
