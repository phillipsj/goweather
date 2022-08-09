package main

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type tickMsg struct{}
type errMsg error

type report struct {
	Loc       string
	Temp      string
	Scale     string
	Condition string
}

type model struct {
	zipcode textinput.Model
	report  report
	err     error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "90210"
	ti.Focus()
	ti.CharLimit = 7
	ti.Width = 20

	return model{
		zipcode: ti,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.zipcode, cmd = m.zipcode.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"What zipcode do you want the weather for (90210)?\n\n%s\n\n%s",
		m.zipcode.View(),
		"(esc to quit)",
	) + "\n"
}

func getWeather(zipcode string) (*report, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?zip=%s,%s&appid=%s")
	return nil, nil
}
