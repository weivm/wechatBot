package pkg

import (
	"context"
	"github.com/eatmoreapple/openwechat"
	"go.uber.org/zap"
	"strings"
	"wechatbot/constant"
	"wechatbot/log"
)

var _ MessageHandler = (*UserMessage)(nil)

type UserMessage struct {
}

func NewUserMessage() *UserMessage {
	return &UserMessage{}
}
func (u *UserMessage) getType() constant.MessageType {
	return constant.UserMessage
}

func (u *UserMessage) receiveMessage(_ context.Context, message *openwechat.Message) string {
	if message.IsText() {
		return message.Content
	}
	return ""
}

func (u *UserMessage) sendMessage(ctx context.Context, message *openwechat.Message, f func(context.Context, string) (string, error)) error {
	var reply = message.Content
	if !message.IsText() {
		reply = "你好！我无法识别"
		_, err := message.ReplyText(reply)
		return err
	}

	reply = strings.TrimSpace(reply)

	reply, err := f(ctx, reply)
	if err != nil {
		log.Logger.Error("sendMessage fail", zap.Error(err))
		return err
	}
	reply = strings.TrimSpace(reply)
	_, err = message.ReplyText(reply)
	if err != nil {
		log.Logger.Error("sendMessage ReplyText fail", zap.Error(err))
	}
	return nil
}
