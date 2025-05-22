package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	active int
	views  []tea.Model
}

func NewModel() *Model {
	return &Model{
		views: []tea.Model{InitFormModel(), NewActive(5)},
	}
}

func (m Model) Init() tea.Cmd {
	return m.views[0].Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.views[m.active], cmd = m.views[m.active].Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case formDoneMsg:
		m.active = 1
		cmds = append(cmds, m.views[m.active].Init())
	}
	return m, cmd
}

func (m Model) View() string {
	if view, ok := m.views[m.active].(tea.Model); ok {
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
