package tui

import (
	"fmt"
	"io"

	"forge/internal/runner"

	tea "github.com/charmbracelet/bubbletea"
)

type progressModel struct {
	event runner.Event
}

var _ tea.Model = progressModel{}

func (m progressModel) Init() tea.Cmd { return nil }
func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case runner.Event:
		m.event = v
	}
	return m, nil
}
func (m progressModel) View() string {
	if m.event.Total == 0 {
		return ""
	}
	return fmt.Sprintf("[%d/%d] %s", m.event.Index, m.event.Total, m.event.Title)
}

func NewProgressObserver(w io.Writer) (func(runner.Event), func()) {
	observe := func(e runner.Event) {
		fmt.Fprintf(w, "%s [%d/%d] %s\n", symbol(e.Type), e.Index, e.Total, e.Title)
		if e.Err != nil {
			fmt.Fprintf(w, "    error: %v\n", e.Err)
		}
	}
	return observe, func() {}
}

func symbol(t runner.EventType) string {
	switch t {
	case runner.EventStarted:
		return "*"
	case runner.EventDone:
		return "+"
	case runner.EventSkipped:
		return "-"
	case runner.EventFailed:
		return "x"
	default:
		return "."
	}
}
