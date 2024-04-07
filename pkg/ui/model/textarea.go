package model

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type TextAreaModel struct {
	textarea textarea.Model
	err      error

  Model
}

func NewTextareModel() TextAreaModel {
	ti := textarea.New()
	ti.Placeholder = "The next thread is about..."
	ti.Focus()

	return TextAreaModel{
		textarea: ti,
		err:      nil,
	}
}

func (m TextAreaModel) Init() tea.Cmd {
  return textarea.Blink
}

func (m TextAreaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  m.textarea, cmd = m.textarea.Update(msg)

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return tea.Quit
    }
  }

  return m, cmd
}

func (m TextAreaModel) View() string { 
  return m.textarea.View()
}

