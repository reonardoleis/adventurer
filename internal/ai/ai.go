package ai

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var cli *openai.Client

func get() *openai.Client {
	if cli == nil {
		cli = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	}

	return cli
}

func Generate(
	prompt string,
	additionalInformation []string,
	maxTokens int,
	temperature float32,
) (string, error) {
	cli := get()
	messages := make([]openai.ChatCompletionMessage, len(additionalInformation)+1)

	for i, info := range additionalInformation {
		messages[i] = openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: info,
		}
	}

	messages[len(additionalInformation)] = openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	}

	resp, err := cli.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo1106,
			MaxTokens:   maxTokens,
			Temperature: temperature,
			Messages:    messages,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func Embed(text string) ([]float32, error) {
	cli := get()
	resp, err := cli.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequestStrings{
			Model: openai.SmallEmbedding3,
			Input: []string{
				text,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Data[0].Embedding, nil
}
