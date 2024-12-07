package main

import (
	"fmt"
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
			handleNewContext(contextName)
		},
	}

	newContextCmd.Flags().StringVar(&contextName, "name", "", "")
	newContextCmd.MarkFlagRequired("name")

	contextCmd.AddCommand(newContextCmd)

	return contextCmd
}

func handleNewContext(name string) {
	fmt.Println(name)
}
