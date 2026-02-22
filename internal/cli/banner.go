package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Banner(version string) string {
	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(lipgloss.Color("63"))
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86"))
	sub := lipgloss.NewStyle().Foreground(lipgloss.Color("246"))
	text := "███████╗ ██████╗ ██████╗  ██████╗ ███████╗\n" +
		"██╔════╝██╔═══██╗██╔══██╗██╔════╝ ██╔════╝\n" +
		"█████╗  ██║   ██║██████╔╝██║  ███╗█████╗\n" +
		"██╔══╝  ██║   ██║██╔══██╗██║   ██║██╔══╝\n" +
		"██║     ╚██████╔╝██║  ██║╚██████╔╝███████╗\n" +
		"╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚══════╝\n\n"
	text += title.Render("Professional Project Scaffolding") + "\n"
	text += sub.Render("Build faster. Ship better.") + "\n"
	text += sub.Render("v" + version)
	return box.Render(text)
}

func CommandMenu() string {
	head := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86")).Render("Available Commands")
	rows := []string{
		"  new         Create a new project interactively",
		"  add         Add a feature to existing project",
		"  doctor      Check your development environment",
		"  config      Manage forge settings",
		"  templates   Browse and manage project templates",
		"  frameworks  List supported frameworks and extras",
		"  version     Show version info",
	}
	return fmt.Sprintf("%s\n%s", head, strings.Join(rows, "\n"))
}
