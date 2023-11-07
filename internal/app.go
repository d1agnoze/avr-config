package internal

import (
	"avr-config/cmd/avr-conf/internal/misc"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Main) Init() tea.Cmd {
	return TickCmd()
}
func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// * Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	// ^ resizing term
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - ProgessBarPadding*2 - 4
		if m.progress.Width > ProgessBarMaxWidth {
			m.progress.Width = ProgessBarMaxWidth
		}
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	// ^ loop the loading screen
	case TickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := m.progress.IncrPercent(4)
		return m, tea.Batch(TickCmd(), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}
func (m Main) View() string {
	// The header
	s := "\nPress q to quit.\n"
	if m.stage == LOADING {
		return misc.LoadingHandler(m.progress, ProgessBarPadding)
	}
	// Send the UI for rendering
	return s
}