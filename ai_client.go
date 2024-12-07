package main

import (
	"os"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type AiClient struct {
	client *anthropic.Client
}

func NewAiClient() *AiClient {
	apiKey := os.Getenv("ANTHROPIC_KEY")
	client := anthropic.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &AiClient{
		client: client,
	}
}

