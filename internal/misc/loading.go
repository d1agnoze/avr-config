package misc

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
)

func LoadingHandler(m progress.Model, padding int) string {
	s := "\nPress q to quit.\n"
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.View() + "\n" +
		pad + s
}
