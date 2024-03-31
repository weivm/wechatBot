package third_party

import (
	"github.com/eatmoreapple/openwechat"
	"go.uber.org/zap"
	"wechatbot/log"
)

func Start(f func(message *openwechat.Message)) {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = f
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("wechat.json")

	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Logger.Error("login fail", zap.Error(err))
			return
		}
	}
	bot.Block()
}
