package main

import (
	"wechatbot/app"
	"wechatbot/configs"
	"wechatbot/log"
)

func main() {
	log.Init()
	configs.InitConfig()
	h := app.InitApp(&configs.Cfg)
	h.Start()
}
