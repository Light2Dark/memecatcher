package internal

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func CreateOpenAIClient() (*openai.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
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
