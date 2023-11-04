package main

import (
	"avr-config/cmd/avr-conf/internal"
	mo "avr-config/cmd/avr-conf/internal/models"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	inputs := []mo.AppInput{
		mo.NewTextField("What's your MCU?"),
		mo.NewTextField("Locate your avrdude.conf location?"),
		mo.NewTextField("What's your Programmer?"),
		mo.NewTextField("Lastly, what's the port your device running on?"),
	}
	model := internal.New(inputs)
	p := tea.NewProgram(*model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Hmmm, there's been an error: %v", err)
		os.Exit(1)
	}
}
