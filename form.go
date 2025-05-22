package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type consumable int

const (
	fruits consumable = iota
	vegetables
	drinks
)

func (c consumable) String() string {
	return []string{"fruit", "vegetable", "drink"}[c]
}

type formModel struct {
	form     *huh.Form
	category consumable
	choice   string
}

type formDoneMsg struct {
	Category consumable
	Choice   string
}

func InitFormModel() formModel {
	var m formModel

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[consumable]().
				Title("What are you in the mood for?").
				Value(&m.category).
				Options(
					huh.NewOption("Some fruit", fruits),
					huh.NewOption("A vegetable", vegetables),
					huh.NewOption("A drink", drinks),
				),

			huh.NewSelect[string]().
				Value(&m.choice).
				Height(7).
				TitleFunc(func() string {
					return fmt.Sprintf("Okay, what kind of %s are you in the mood for?", m.category)
				}, &m.category).
				OptionsFunc(func() []huh.Option[string] {
					switch m.category {
					case fruits:
						return []huh.Option[string]{
							huh.NewOption("Tangerine", "tangerine"),
							huh.NewOption("Canteloupe", "canteloupe"),
							huh.NewOption("Pomelo", "pomelo"),
							huh.NewOption("Grapefruit", "grapefruit"),
						}
					case vegetables:
						return []huh.Option[string]{
							huh.NewOption("Carrot", "carrot"),
							huh.NewOption("Jicama", "jicama"),
							huh.NewOption("Kohlrabi", "kohlrabi"),
							huh.NewOption("Fennel", "fennel"),
							huh.NewOption("Ginger", "ginger"),
						}
					default:
						return []huh.Option[string]{
							huh.NewOption("Coffee", "coffee"),
							huh.NewOption("Tea", "tea"),
							huh.NewOption("Bubble Tea", "bubble tea"),
							huh.NewOption("Agua Fresca", "agua-fresca"),
						}
					}
				}, &m.category),
		),
	)

	m.form = form
	return m
}

func (m formModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	form, cmd := m.form.Update(msg)
	m.form = any(form).(*huh.Form)

	if m.form.State == huh.StateCompleted {
		return m, func() tea.Msg {
			return formDoneMsg{
				Category: m.category,
				Choice:   m.choice,
			}
		}
	}

	return m, cmd
}

func (m formModel) View() string {
	return m.form.View()
}
