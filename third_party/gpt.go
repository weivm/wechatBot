package third_party

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"wechatbot/configs"
	"wechatbot/constant"
)

type AiChatClient struct {
	client *openai.Client
	cfg    *configs.Config
}

func NewAiChatClient(cfg *configs.Config) *AiChatClient {
	var c *openai.Client
	if cfg.OpenaiType == constant.OpenAi {
		c = openai.NewClient(cfg.AccessKey)
	} else {
		defaultConfig := openai.DefaultAzureConfig(cfg.AccessKey, cfg.BaseUrl)
		defaultConfig.AzureModelMapperFunc = func(model string) string {
			return cfg.EndPoint
		}
		c = openai.NewClientWithConfig(defaultConfig)
	}
	return &AiChatClient{client: c, cfg: cfg}
}

// todo
func (n *AiChatClient) getPrompt() string {
	return ""
}

func (n *AiChatClient) getType() constant.AiChatType {
	return constant.GptChat
}

func (n *AiChatClient) Reply(ctx context.Context, q string) (string, error) {
	messages := userQuestionMessage(n.getPrompt(), q)
	rsp, err := n.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:            n.cfg.Model,
			Messages:         messages,
			Temperature:      0,
			TopP:             0.2,
			FrequencyPenalty: 1,
		})
	if err != nil {
		return "", err
	}

	if len(rsp.Choices) > 0 {
		return rsp.Choices[0].Message.Content, nil
	}
	return "", nil
}
