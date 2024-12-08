package main

import (
	"iacoder/pkg/core"
	"iacoder/pkg/ui"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")

	coder := core.NewCoder()

	errChan := make(chan error)

	chat := core.NewChat(coder, errChan)
	app := ui.NewApp(chat)

	go func() {
		for err := range errChan {
			if err != nil {
				app.GetProgram().Quit()
			}
		}
	}()

	app.Start()
}

