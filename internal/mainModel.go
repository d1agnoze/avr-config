package internal

import (
	"avr-config/cmd/avr-conf/internal/misc"
	"avr-config/cmd/avr-conf/internal/models"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	LOADING    = 1
	SELECTION  = 2
	GENERATING = 3
	FINISH     = 4
)
const (
	ProgessBarPadding  = 2
	ProgessBarMaxWidth = 80
)

// * Main begin here!!!
type Main struct {
	progress   progress.Model
	stage      int //identifier for stages of the application
	inputs     []models.AppInput
	width      int
	height     int
	index      int
	radio      models.Radio
	keyParams  misc.Template
	scanResult map[string]string
}

func New(input []models.AppInput) *Main {
	return &Main{
		inputs:   input,
		stage:    LOADING,
		progress: progress.New(progress.WithGradient("#11998e", "#38ef7d")),
	}
}

// ^ to set time for the loader
type TickMsg time.Time

func TickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
