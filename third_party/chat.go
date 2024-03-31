package third_party

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"strings"
	"wechatbot/configs"
	"wechatbot/constant"
)

type ChatCompletion interface {
	getPrompt() string // getPrompt returns the prompt for the chat completion.
	Reply(ctx context.Context, q string) (string, error)
	getType() constant.AiChatType
}

type AiChat struct {
	Client map[constant.AiChatType]ChatCompletion
}

func NewAiChat(cfg *configs.Config) *AiChat {
	n := &AiChat{}
	n.registerAiChat(NewAiChatClient(cfg), NewCusAiChatClient(cfg))
	return n
}
func (a *AiChat) registerAiChat(c ...ChatCompletion) {
	var clients = make(map[constant.AiChatType]ChatCompletion)
	for _, v := range c {
		clients[v.getType()] = v
	}
	a.Client = clients
}

func baseQuestionMessage(prompt string) []openai.ChatCompletionMessage {
	var messages []openai.ChatCompletionMessage
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: " You are a versatile problem-solving assistant ",
	})

	if prompt != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: prompt,
		})
	}
	return messages
}

func userQuestionMessage(prompt, question string) []openai.ChatCompletionMessage {
	messages := baseQuestionMessage(prompt)
	q := strings.TrimSpace(question)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: q,
	})
	return messages
}
