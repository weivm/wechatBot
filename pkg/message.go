package pkg

import (
	"context"
	"github.com/eatmoreapple/openwechat"
	"go.uber.org/zap"
	"wechatbot/configs"
	"wechatbot/constant"
	"wechatbot/log"
	"wechatbot/third_party"
)

type MessageHandler interface {
	getType() constant.MessageType
	receiveMessage(ctx context.Context, message *openwechat.Message) string
	sendMessage(ctx context.Context, message *openwechat.Message, f func(context.Context, string) (string, error)) error
}

type MessagesHandle struct {
	*third_party.AiChat
	cfg            *configs.Config
	messageHandles map[constant.MessageType]MessageHandler
}

func NewHandle(a *third_party.AiChat, cfg *configs.Config) *MessagesHandle {
	var m = &MessagesHandle{AiChat: a, cfg: cfg}
	m.RegisterHandle(NewGroupMessage(), NewUserMessage())
	return m
}

func (m *MessagesHandle) Start() {
	third_party.Start(m.HandleMessage)
}
func (m *MessagesHandle) RegisterHandle(hd ...MessageHandler) {
	var h = map[constant.MessageType]MessageHandler{}
	for _, v := range hd {
		h[v.getType()] = v
	}
	m.messageHandles = h
}
func (m *MessagesHandle) HandleMessage(message *openwechat.Message) {
	typ := m.getMessageType(message)
	ctx := context.Background()
	m.getMessageHandle()(ctx, message, typ)
	return
}

func (m *MessagesHandle) send(ctx context.Context, message string) (string, error) {
	chatClient := m.Client[constant.CusGptChat]
	reply, err := chatClient.Reply(ctx, message)
	if err != nil {
		log.Logger.Error("getMessageHandle Reply fail", zap.Error(err))
		return "", err
	}
	if reply == "" {
		reply = m.cfg.DefaultReply
	}

	return reply, nil
}
func (m *MessagesHandle) getMessageType(message *openwechat.Message) constant.WechatMessageType {
	if message.IsSendByGroup() {
		return constant.WechatGroupMessage
	}

	if message.IsFriendAdd() {
		return constant.WechatAddFriendMessage
	}

	if message.IsSendByFriend() {
		return constant.WechatUserMessage
	}

	if message.IsJoinGroup() {
		return constant.WechatJoinGroupMessage
	}
	return 0
}

func (m *MessagesHandle) getMessageHandle() func(ctx context.Context, message *openwechat.Message, typ constant.WechatMessageType) error {
	return func(ctx context.Context, message *openwechat.Message, typ constant.WechatMessageType) error {

		switch typ {
		case constant.WechatGroupMessage:
			err := m.messageHandles[constant.GroupMessage].sendMessage(ctx, message, m.send)
			if err != nil {
				return err
			}
		case constant.WechatAddFriendMessage:
			_, err := message.Agree("你好，请问需要我提供什么帮助？")
			if err != nil {
				log.Logger.Error("getMessageHandle agree fail", zap.Error(err))
				return err
			}

		case constant.WechatUserMessage:
			err := m.messageHandles[constant.UserMessage].sendMessage(ctx, message, m.send)
			if err != nil {
				return err
			}

		case constant.WechatJoinGroupMessage:
			_, err := message.Agree("欢迎进群")
			if err != nil {
				log.Logger.Error("getMessageHandle agree fail", zap.Error(err))
				return err
			}
		}
		return nil
	}
}
