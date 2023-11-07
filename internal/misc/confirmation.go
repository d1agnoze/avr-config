package misc

import (
	"fmt"
	"log"

	"github.com/charmbracelet/lipgloss"
)

func Confirmation(reps Template) string {
	log.Println(reps)
	recap := fmt.Sprintf("Confirm all your input?\nMCU - %s\nProgrammer - %s\nPort - %s\n", reps.MCU, reps.PROGRAMER, reps.PORT_NAME)
	return lipgloss.JoinVertical(lipgloss.Left, CC, recap)
}
