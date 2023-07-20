package api

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetData(query string) string {

	token := os.Getenv("Token")
	client := openai.NewClient(token)
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		log.Printf("error occurred here:  %v", err)
		return "error"
	}

	return response.Choices[0].Message.Content
}
