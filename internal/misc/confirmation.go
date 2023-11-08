package misc

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Confirmation(reps Template) string {
	recap := fmt.Sprintf("Confirm all your input?\nMCU - %s\nProgrammer - %s\nPort - %s\n", reps.MCU, reps.PROGRAMER, reps.PORT_NAME)
	return lipgloss.JoinVertical(lipgloss.Left, CC, recap)
}
