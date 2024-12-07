package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type RateLimits struct {
	RequestsLimit     int
	RequestsRemaining int
	RequestsReset     time.Time

	InputTokensLimit     int
	InputTokensRemaining int
	InputTokensReset     time.Time

	OutputTokensLimit     int
	OutputTokensRemaining int
	OutputTokensReset     time.Time

	RetryAfter int
}

type UsageStats struct {
	RequestsTotal int
	//RequestsInLastMinute int

	InputTokensTotal int
	//InputTokensInLastMinute int

	OutputTokensTotal int
	//OutputTokensInLastMinute int
}

type AiClient struct {
	client *anthropic.Client
	Limits *RateLimits
	Usage  UsageStats
}

func NewAiClient() *AiClient {
	apiKey := os.Getenv("ANTHROPIC_KEY")
	client := anthropic.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &AiClient{
		client: client,
		Limits: nil,
		Usage: UsageStats{
			RequestsTotal: 0,
			InputTokensTotal: 0,
			OutputTokensTotal: 0,
		},
	}
}

func (c *AiClient) SendMessage(input string) ([]string, error) {
	systemPrompt, err := os.ReadFile("system.prompt")
	if err != nil {
		return nil, err
	}

	response, err := c.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model: anthropic.F(anthropic.ModelClaude3_5SonnetLatest),
		MaxTokens: anthropic.Int(4096),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(fmt.Sprintf("<Task>%s</Task>", input))),
		}),
		System: anthropic.F([]anthropic.TextBlockParam{
			anthropic.NewTextBlock(string(systemPrompt)),
		}),
	})
	if err != nil {
		return nil, err
	}

	output := []string{}
	for _, block := range response.Content {
		if block.Type != anthropic.ContentBlockTypeText {
			continue
		}
		output = append(output, block.Text)
	}
	return output, nil
}

