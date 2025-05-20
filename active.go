package main

import (
	"time"

	"github.com/charmbracelet/bubbles/v2/progress"
	"github.com/charmbracelet/bubbles/v2/timer"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss/v2"
)

type active struct {
	secsLeft float64
	title    string
	timer    timer.Model
	progress progress.Model
	form     *huh.Form
}

func NewActive(secs float64) active {
	// create a timer with the desired interval
	return active{
		title:    "Get to work!",
		secsLeft: secs,
		timer:    timer.New(time.Second * time.Duration(secs)),
		progress: progress.New(
			progress.WithDefaultGradient(),
			progress.WithWidth(40),
		),
		form: huh.NewForm(huh.NewGroup(
			huh.NewInput(),
		)),
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

	if (float64(m.secsLeft)-m.timer.Timeout.Seconds())/float64(m.secsLeft) == 1 {
		m.title = "Done! Take a break!"
	}
	return m, tea.Batch(cmds...)
}

func (m active) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		m.title,
		m.progress.ViewAs(float64(m.secsLeft-m.timer.Timeout.Seconds())/m.secsLeft),
		m.timer.View(),
	)
}
