package third_party

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"io/ioutil"
	"net/http"
	"wechatbot/configs"
	"wechatbot/constant"
)

type CusAiChatClient struct {
	client *openai.Client
}

func NewCusAiChatClient(cfg *configs.Config) *CusAiChatClient {
	c := openai.NewClient(cfg.AccessKey)
	return &CusAiChatClient{client: c}
}

// todo
func (n *CusAiChatClient) getPrompt() string {
	return ""
}

func (n *CusAiChatClient) getType() constant.AiChatType {
	return constant.CusGptChat
}

func (n *CusAiChatClient) Reply(ctx context.Context, q string) (string, error) {
	messages := userQuestionMessage(n.getPrompt(), q)
	param := openai.ChatCompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Messages:         messages,
		Temperature:      0,
		TopP:             0.2,
		FrequencyPenalty: 1,
	}
	requestData, err := json.Marshal(param)

	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", configs.Cfg.BaseUrl+"chat/completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+configs.Cfg.AccessKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	result := &openai.ChatCompletionResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return "", err
	}

	var reply string
	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}
	return reply, nil
}
