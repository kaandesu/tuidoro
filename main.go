package main

import (
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
)

type Model struct {
	active    int
	views     []tea.Model
	debugDump io.Writer
}

func NewModel(dump io.Writer) *Model {
	return &Model{
		views:     []tea.Model{InitFormModel(), NewActive(5)},
		debugDump: dump,
		active:    0,
	}
}

func (m Model) Init() tea.Cmd {
	return m.views[0].Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.views[m.active], cmd = m.views[m.active].Update(msg)
	if m.debugDump != nil {
		spew.Fdump(m.debugDump, msg)
	}
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
	var dump *os.File
	var err error

	if _, ok := os.LookupEnv("DEBUG"); ok {
		dump, err = os.OpenFile("messages.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}

	p := tea.NewProgram(NewModel(dump))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
