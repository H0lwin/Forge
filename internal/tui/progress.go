package tui

import (
	"fmt"
	"io"
	"strings"
	"time"

	"forge/internal/runner"
)

func NewProgressObserver(w io.Writer) (func(runner.Event), func()) {
	started := time.Now()
	observe := func(e runner.Event) {
		sym := symbol(e.Type)
		bar := progressBar(e.Index, e.Total, 20)
		pct := progressPercent(e.Index, e.Total)
		fmt.Fprintf(w, "%s [%2d/%2d] %3d%% %s %s", sym, e.Index, e.Total, pct, bar, e.Title)
		if e.Type == runner.EventDone || e.Type == runner.EventFailed {
			fmt.Fprintf(w, "  (%s)", time.Since(started).Round(time.Second))
		}
		fmt.Fprintln(w)
		if e.Err != nil {
			fmt.Fprintf(w, "    error: %v\n", e.Err)
		}
	}
	return observe, func() {}
}

func symbol(t runner.EventType) string {
	switch t {
	case runner.EventStarted:
		return ">"
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

func progressPercent(index, total int) int {
	if total <= 0 {
		return 0
	}
	p := int(float64(index) / float64(total) * 100.0)
	if p > 100 {
		return 100
	}
	if p < 0 {
		return 0
	}
	return p
}

func progressBar(index, total, width int) string {
	if width <= 0 {
		return ""
	}
	if total <= 0 {
		return "[" + strings.Repeat("-", width) + "]"
	}
	filled := int(float64(index) / float64(total) * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	return "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + "]"
}
