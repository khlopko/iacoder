package ui

import (
	"fmt"
	"iacoder/pkg/core"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	//"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
)

type ModelHandlers interface {
	HandleEnter(input string)
}

type Model struct {
	viewport viewport.Model
	textarea textarea.Model
	spinner  spinner.Model

	handlers ModelHandlers

	isLoading bool
}

func NewModel() *Model {
	vp := viewport.New(1, 5)
	vp.SetContent("What can I do for you?")

	ta := textarea.New()
	ta.Placeholder = "Write your task..."
	ta.Prompt = "\t> "
	ta.ShowLineNumbers = false
	ta.SetHeight(5)
	ta.Focus()

	s := spinner.New()
	s.Spinner = spinner.Pulse
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))

	return &Model{
		viewport:  vp,
		textarea:  ta,
		spinner:   s,
		handlers:  nil,
		isLoading: false,
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, m.spinner.Tick)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case error:
		m.isLoading = false
		m.viewport.SetContent(msg.Error())

	case []core.TaskResult:
		m.isLoading = false
		content := ""
		for _, result := range msg {
			if result.Err != nil {
				content += result.Err.Error()
				content += "\n"
			} else if result.Message != nil {
				content += *result.Message
				content += "\n"
			}
		}
		m.viewport.SetContent(content)

	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 5
		m.textarea.SetWidth(msg.Width)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "ctrl+c":
		return m, tea.Quit
	case "enter":
		input := m.textarea.Value()
		go func() {
			m.handlers.HandleEnter(input)
		}()

		m.textarea.Reset()
		m.isLoading = true
		return m, m.spinner.Tick
	default:
		var cmd tea.Cmd
		m.textarea, cmd = m.textarea.Update(msg)
		return m, cmd
	}
}

func (m *Model) View() string {
	if m.isLoading {
		return m.spinner.View()
	}
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(), m.textarea.View(),
	)
}
