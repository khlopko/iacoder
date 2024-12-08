package ui

import (
	"fmt"
	"iacoder/pkg/core"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	chat    *core.Chat
	program *tea.Program
}

func NewApp(chat *core.Chat) *App {
	model := NewModel()
	program := tea.NewProgram(model, tea.WithAltScreen())
	app := App{chat, program}
	model.handlers = &app
	return &app
}

func (self *App) Start() {
	_, err := self.program.Run()
	if err != nil {
		fmt.Printf("Failed to run app with error: %+v\n", err)
		os.Exit(1)
	}
}

func (self *App) HandleEnter(input string) {
	result, err := self.chat.ExecuteNewTask(input)
	if err != nil {
		self.program.Send(err)
	} else {
		self.program.Send(result)
	}
}
