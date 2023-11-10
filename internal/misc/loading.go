package misc

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

func LoadingHandler(m progress.Model, padding int) string {
	aa := BB
	var asciiStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6fc452")).MarginTop(0)
	pad := strings.Repeat(" ", padding)
	return asciiStyle.Render(aa) + "\n" +
		pad + m.View() + "\n" +
		pad
}
