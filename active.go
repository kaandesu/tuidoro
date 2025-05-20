package main

import (
	"time"

	"github.com/charmbracelet/bubbles/v2/list"
	"github.com/charmbracelet/bubbles/v2/progress"
	"github.com/charmbracelet/bubbles/v2/timer"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type active struct {
	timer    timer.Model
	list     list.Model
	progress progress.Model
}

func NewActive() active {
	// create a timer with the desired interval
	return active{
		timer:    timer.New(time.Second * 10),
		progress: progress.New(),
	}
}

func (m active) Init() tea.Cmd {
	// return tea.Batch(m.timer.Init(), m.progress.Init())
	return m.timer.Init()
}

func (m active) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.timer, cmd = m.timer.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m active) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		"Get to work",
		m.progress.ViewAs(float64(10-m.timer.Timeout.Seconds())/10.0),
		m.timer.View(),
	)
}
