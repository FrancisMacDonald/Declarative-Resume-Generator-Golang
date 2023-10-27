package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Initialize(token string) {
	client := openai.NewClient(token)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello OpenAi!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("Error running ChatCompletion: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
