package misc

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

func LoadingHandler(m progress.Model, padding int) string {
	s := "\nPress q to quit.\n"
	// aa := `
	// _______  __   __  ______           _______  _______  __    _  _______
	// |   _   ||  | |  ||    _ |         |       ||       ||  |  | ||       |
	// |  |_|  ||  |_|  ||   | ||   ____  |       ||   _   ||   |_| ||    ___|
	// |       ||       ||   |_||_ |____| |       ||  | |  ||       ||   |___
	// |       ||       ||    __  |       |      _||  |_|  ||  _    ||    ___|
	// |   _   | |     | |   |  | |       |     |_ |       || | |   ||   |
	// |__| |__|  |___|  |___|  |_|       |_______||_______||_|  |__||___|
	//  `
	aa:=`
	█████╗ ██╗   ██╗██████╗        ██████╗ ██████╗ ███╗   ██╗███████╗██╗ ██████╗
	██╔══██╗██║   ██║██╔══██╗      ██╔════╝██╔═══██╗████╗  ██║██╔════╝██║██╔════╝
	███████║██║   ██║██████╔╝█████╗██║     ██║   ██║██╔██╗ ██║█████╗  ██║██║  ███╗
	██╔══██║╚██╗ ██╔╝██╔══██╗╚════╝██║     ██║   ██║██║╚██╗██║██╔══╝  ██║██║   ██║
	██║  ██║ ╚████╔╝ ██║  ██║      ╚██████╗╚██████╔╝██║ ╚████║██║     ██║╚██████╔╝
	╚═╝  ╚═╝  ╚═══╝  ╚═╝  ╚═╝       ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝     ╚═╝ ╚═════╝
	`
	// aa := `
	
	// ░█▀█░█░█░█▀▄░░░░░█▀▀░█▀█░█▀█░█▀▀░▀█▀░█▀▀
	// ░█▀█░▀▄▀░█▀▄░▄▄▄░█░░░█░█░█░█░█▀▀░░█░░█░█
	// ░▀░▀░░▀░░▀░▀░░░░░▀▀▀░▀▀▀░▀░▀░▀░░░▀▀▀░▀▀▀
	// `
	var asciiStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6fc452")).MarginTop(0)
	pad := strings.Repeat(" ", padding)
	return asciiStyle.Render(aa) + "\n" +
		pad + m.View() + "\n" +
		pad + s
}
