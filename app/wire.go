//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"wechatbot/configs"
	"wechatbot/pkg"
	"wechatbot/third_party"
)

func InitApp(cfg *configs.Config) *pkg.MessagesHandle {
	panic(wire.Build(third_party.NewAiChat, pkg.NewHandle))
}
