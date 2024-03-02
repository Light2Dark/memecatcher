package internal

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func CreateOpenAIClient(openaiKey string) (*openai.Client, error) {
	client := openai.NewClient(openaiKey)
	return client, nil
}

func ChatCompletion(client *openai.Client, prompt string, maxTokens int) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: maxTokens,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
