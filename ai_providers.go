package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strings"
)

type AiProvider interface {
	Initialize(token string)
	CheckSpellingGrammar(text string) CorrectedText
}

type OpenAiProvider struct {
	client *openai.Client
}

func (provider OpenAiProvider) Initialize(token string) {
	client := openai.NewClient(token)

	// test connection
	_, err := client.ListEngines(context.Background())

	if err != nil {
		fmt.Printf("Error listing engines: %v\n", err)
	}

	provider.client = client
}

func (provider OpenAiProvider) CheckSpellingGrammar(text string) CorrectedText {
	// TODO: Allow individual corrections.
	// TODO: Return a reason for the corrections.

	if strings.TrimSpace(text) == "" {
		return CorrectedText{
			Original:  text,
			Corrected: text,
			Changes:   "",
		}
	}

	resp, err := provider.client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Prompt: "Correct any spelling and grammar mistakes in the text. Return the full text only. Maintain any newlines.\n" +
				"Original text: " + text + "\n" +
				"Corrected text:",
			MaxTokens:   64,
			Temperature: 0.7,
			TopP:        1,
			N:           1,
			Stream:      false,
			LogProbs:    0,
			Stop:        nil,
		},
	)

	if err != nil {
		fmt.Printf("Error running Completion: %v\n", err)
	}

	fmt.Println(resp.Choices[0].Text)

	correctedText := resp.Choices[0].Text
	correctedChanges := "" // TODO: add reason for corrections

	return CorrectedText{
		Original:  text,
		Corrected: correctedText,
		Changes:   correctedChanges,
	}
}
