package models

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Blink() tea.Msg
	Blur() tea.Msg
	Focus() tea.Cmd
	SetValue(string)
	Value() string
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

// We need to have a wrapper for our bubbles as they don't currently implement the tea.Model interface
// textinput, textarea

type shortAnswerField struct {
	textinput textinput.Model
}

func newShortAnswerField() *shortAnswerField {
	a := shortAnswerField{}
	model := textinput.New()
	model.Placeholder = "Type here..."
	model.Focus()
	a.textinput = model
	return &a
}

func (a *shortAnswerField) Blink() tea.Msg {
	return textinput.Blink()
}

func (a *shortAnswerField) Init() tea.Cmd {
	return nil
}

func (a *shortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	a.textinput, cmd = a.textinput.Update(msg)
	return a, cmd
}

func (a *shortAnswerField) View() string {
	return a.textinput.View()
}

func (a *shortAnswerField) Focus() tea.Cmd {
	return a.textinput.Focus()
}

func (a *shortAnswerField) SetValue(s string) {
	a.textinput.SetValue(s)
}

func (a *shortAnswerField) Blur() tea.Msg {
	return a.textinput.Blur
}

func (a *shortAnswerField) Value() string {
	return a.textinput.Value()
}
