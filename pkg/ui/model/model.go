package model

import tea "github.com/charmbracelet/bubbletea"

type Model interface {
  Update(msg tea.Msg) (tea.Model, tea.Cmd)
  View() string
  Init() tea.Cmd

  tea.Model
}
