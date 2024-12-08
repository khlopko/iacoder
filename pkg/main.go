package main

import (
	"fmt"
	"iacoder/pkg/core"
	"iacoder/pkg/ui"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")

	coder := core.NewCoder()

	chat, err := core.NewChat(coder)
	if err != nil {
		fmt.Printf("Failed to create core systems: %+v", err)
		os.Exit(1)
	}

	app := ui.NewApp(chat)
	app.Start()
}

