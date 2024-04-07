package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/filipgorny/potion/pkg/ui/model"
)

type Ui struct {
  program *tea.Program
}

func NewUI(model model.Model) *Ui {
  return &Ui{program: tea.NewProgram(model, tea.WithAltScreen())}
}

func (u *Ui) Run() {
  error, _ := u.program.Run()

  if error != nil {
    panic(error)
  }

}
