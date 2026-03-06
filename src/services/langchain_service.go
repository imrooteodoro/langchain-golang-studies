package services

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func CreateOllamaClient(modelName string) (*ollama.LLM, error) {
	client, err := ollama.New(ollama.WithModel(modelName))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ChatWithOllama(ctx context.Context, client *ollama.LLM, prompt string) (string, error) {
	response, err := llms.GenerateFromSinglePrompt(ctx, client, prompt)
	if err != nil {
		return "", err
	}
	return response, nil
}