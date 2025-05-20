package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type Model struct {
	// active, new, history?, break?
	active int
	views  []tea.Model
}

func NewModel() *Model {
	return &Model{
		views: []tea.Model{NewActive(5)},
	}
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, view := range m.views {
		cmds = append(cmds, view.Init())
	}

	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.views[m.active], cmd = m.views[m.active].Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m Model) View() string {
	if view, ok := m.views[m.active].(tea.ViewModel); ok {
		return view.View()
	}
	return "no view model :("
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
