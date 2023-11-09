package main

import (
    "context"
    "fmt"
    "github.com/sashabaranov/go-openai"
    "math"
    "strings"
)

type AiProvider interface {
    Initialize(token string, initialPrompt string, seed *int)
    CheckSpellingGrammar(text string) CorrectedText
}

type OpenAiProvider struct {
    client        *openai.Client
    seed          *int
    initialPrompt string // You are a helpful blah blah blah
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

    chatCompletionRequest := openai.ChatCompletionRequest{
        Model:       openai.GPT3Dot5Turbo,
        Temperature: math.SmallestNonzeroFloat32, // Testing if this helps to make it deterministic.
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleSystem,
                Content: provider.initialPrompt,
            },
            {
                Role:    openai.ChatMessageRoleUser,
                Content: text,
            },
        },
    }

    // If there is a seed on the provider, use it.
    if provider.seed != nil {
        chatCompletionRequest.Seed = provider.seed
    }

    resp, err := provider.client.CreateChatCompletion(
        context.Background(),
        chatCompletionRequest,
    )

    /*
       	Chat Completions are non-deterministic by default (which means model outputs may differ from request to request). That being said, we offer some control towards deterministic outputs by giving you access to the seed parameter and the system_fingerprint response field.

              To receive (mostly) deterministic outputs across API calls, you can:

                  Set the seed parameter to any integer of your choice and use the same value across requests you'd like deterministic outputs for.
                  Ensure all other parameters (like prompt or temperature) are the exact same across requests.

              Sometimes, determinism may be impacted due to necessary changes OpenAI makes to model configurations on our end. To help you keep track of these changes, we expose the system_fingerprint field. If this value is different, you may see different outputs due to changes we've made on our systems.
    */

    if err != nil {
        fmt.Printf("ChatCompletion error: %v\n", err)
        return CorrectedText{}
    }

    messageContent := resp.Choices[0].Message.Content
    // systemFingerprint := resp.SystemFingerprint // Not implemented in the library yet.
    fmt.Println(messageContent)

    correctedText := messageContent
    correctedChanges := "" // TODO: add reason for corrections

    return CorrectedText{
        Original:  text,
        Corrected: correctedText,
        Changes:   correctedChanges,
    }
}
