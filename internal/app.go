package internal

import (
	"avr-config/cmd/avr-conf/internal/misc"
	"avr-config/cmd/avr-conf/internal/models"
	"sync"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var wg sync.WaitGroup

func (m Main) Init() tea.Cmd {
	wg.Add(1)
	go misc.Scan(&wg)
	return TickCmd()
}
func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	current := m.inputs[m.index]
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// * Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "esc":
			return m, tea.Quit

		case "enter":
			if m.stage == SELECTION {
				current.Value = current.Input.Value()
				m.inputs[m.index].Value = current.Input.Value()
				if m.index == len(m.inputs)-1 {
					m.keyParams = templateMapper(m)
					m.stage = GENERATING
				}
				m.Next()
				return m, current.Input.Blur
			}
			if m.stage == FINISH {
				return m, tea.Quit
			}
			if m.stage == GENERATING {
				m.radio.Choice = models.Bool_choices[m.radio.Cursor]
			}
			return m, nil
		case "down", "j":
			if m.stage == GENERATING {
				m.radio.Cursor++
				if m.radio.Cursor >= len(models.Bool_choices) {
					m.radio.Cursor = 0
				}
			}

		case "up", "k":
			if m.stage == GENERATING {
				m.radio.Cursor--
				if m.radio.Cursor < 0 {
					m.radio.Cursor = len(models.Bool_choices) - 1
				}
			}
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
			wg.Wait()
			m.stage = SELECTION
			return m, nil
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
	current.Input, cmd = current.Input.Update(msg)
	return m, cmd
}
func (m Main) View() string {
	current := m.inputs[m.index]
	// The header
	s := "\nPress ESC or to quit.\n"
	switch m.stage {
	case LOADING:
		s += misc.LoadingHandler(m.progress, ProgessBarPadding)
	case SELECTION:
		s += lipgloss.JoinVertical(lipgloss.Left, misc.CC, current.Question, current.Input.View())
	case GENERATING:
		s += lipgloss.JoinVertical(lipgloss.Left, misc.Confirmation(m.keyParams), m.radio.View())
	}
	return s
}
func (m *Main) Next() {
	if m.index < len(m.inputs)-1 {
		m.index++
	} else {
		m.index = 0
	}
}
func templateMapper(m Main) misc.Template {
	res := misc.NewTemplate()
	res.MCU = m.inputs[0].Value
	res.PROGRAMER = m.inputs[1].Value
	res.PORT_NAME = m.inputs[2].Value
	return res
}
