package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")

	aiClient := NewAiClient()
	cliBuilder := NewCliBuilder(aiClient)
	rootCmd := cliBuilder.Build()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
