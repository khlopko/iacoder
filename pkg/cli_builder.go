package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CliBuilder struct {
	aiClient *AiClient
}

func NewCliBuilder(aiClient *AiClient) *CliBuilder {
	return &CliBuilder{
		aiClient: aiClient,
	}
}

func (b *CliBuilder) Build() *cobra.Command {
	rootCmd := &cobra.Command {
		Use: "iacoder",
		Short: "I am an ultimate AI helper from CLI",
	}

	rootCmd.AddCommand(b.makeContextCommand())
	return rootCmd
}

func (b *CliBuilder) makeContextCommand() *cobra.Command {
	contextCmd := &cobra.Command{
		Use: "context",
	}

	var contextName string

	newContextCmd := &cobra.Command {
		Use: "new",
		Run: func(cmd *cobra.Command, args []string) {
			b.handleNewContext(contextName)
		},
	}

	newContextCmd.Flags().StringVar(&contextName, "name", "", "")
	newContextCmd.MarkFlagRequired("name")
	contextCmd.AddCommand(newContextCmd)

	demoContextCmd := &cobra.Command {
		Use: "demo",
		Run: func(cmd *cobra.Command, args []string) {
			b.handleDemoContext(contextName)
		},
	}

	demoContextCmd.Flags().StringVar(&contextName, "task", "", "")
	demoContextCmd.MarkFlagRequired("task")
	contextCmd.AddCommand(demoContextCmd)

	return contextCmd
}

func (b *CliBuilder) handleNewContext(name string) {
	fmt.Println(name)
}

type ParsedMessage struct {
	Path string `json:"path"`
	Content string `json:"content"`
}

func (b *CliBuilder) handleDemoContext(task string) {
	messages, err := b.aiClient.SendMessage(task)
	if err != nil {
		fmt.Printf("Couldn't get anything from AI with error: %+v\n", err)
		return
	}

	if len(messages) == 0 {
		fmt.Println("AI didn't provide any response :(")
	}

	for _, message := range messages {
		var parsedMessages []ParsedMessage
		err := json.Unmarshal([]byte(message), &parsedMessages)
		if err != nil {
			fmt.Printf("Failed to decode due to error: %+v\n", err)
			fmt.Printf("Original message is: %s\n", message)
		}

		for _, parsedMessage := range parsedMessages {
			err = os.WriteFile(parsedMessage.Path, []byte(parsedMessage.Content), os.ModeAppend)
			if err != nil {
				fmt.Printf("Failed to write at %s due to error: %+v\n", parsedMessage.Path, err)
			} else {
				fmt.Printf("Did write at %s\n", parsedMessage.Path)
			}
		}
	}

	if err == nil {
		fmt.Println("All done!")
	}
}
