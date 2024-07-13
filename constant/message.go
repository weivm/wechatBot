package constant

type MessageType int

const (
	UserMessage MessageType = iota + 1
	GroupMessage
)

type AiChatType int

const (
	GptChat = iota + 1
	CusGptChat
)

type WechatMessageType int

const (
	WechatGroupMessage WechatMessageType = iota + 1
	WechatUserMessage
	WechatAddFriendMessage
	WechatJoinGroupMessage
)

// openai type
const (
	OpenAi = "openai"
	Azure  = "azure"
)
