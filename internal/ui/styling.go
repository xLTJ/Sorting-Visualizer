package ui

import "github.com/charmbracelet/lipgloss"

var (
	AppStyle = lipgloss.NewStyle().
		Padding(1, 2)

	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#c27aff")).
		Padding(0, 2)

	StatsStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#c27aff")).
		Padding(1, 2).
		Width(30).
		BorderTop(false)

	HelpStyle = lipgloss.NewStyle().
		PaddingLeft(4).
		PaddingBottom(1)

	WindowStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#c27aff")).
		Padding(0, 1).
		BorderTop(false)

	WindowTitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c27aff"))

	AxisStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("5"))

	BlockStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fff"))

	ActiveBlockStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff6467"))
)
