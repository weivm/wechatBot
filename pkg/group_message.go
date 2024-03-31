package pkg

import (
	"context"
	"github.com/eatmoreapple/openwechat"
	"go.uber.org/zap"
	"strings"
	"wechatbot/constant"
	"wechatbot/log"
)

var _ MessageHandler = (*GroupMessage)(nil)

type GroupMessage struct {
}

func NewGroupMessage() *GroupMessage {
	return &GroupMessage{}
}
func (g *GroupMessage) getType() constant.MessageType {
	return constant.GroupMessage
}

func (g *GroupMessage) receiveMessage(_ context.Context, message *openwechat.Message) string {
	if message.IsText() {
		return message.Content
	}
	return ""
}

func (g *GroupMessage) sendMessage(ctx context.Context, message *openwechat.Message, f func(context.Context, string) (string, error)) error {
	if !message.IsAt() {
		return nil
	}

	sender, _ := message.Sender()
	atSenderName := "@" + sender.NickName
	content := strings.ReplaceAll(message.Content, atSenderName, "")
	content = strings.TrimSpace(content)

	reply, err := f(ctx, content)
	if err != nil {
		log.Logger.Error("sendMessage fail", zap.Error(err))
		return err
	}

	reply = strings.TrimSpace(reply)

	groupSender, err := message.SenderInGroup()
	if err != nil {
		log.Logger.Error("getMessageHandle SenderInGroup fail", zap.Error(err))
		return err
	}

	atSenderName = "@" + groupSender.NickName
	reply = atSenderName + reply

	_, err = message.ReplyText(reply)
	if err != nil {
		log.Logger.Error("getMessageHandle ReplyText fail", zap.Error(err))
	}
	return nil
}
