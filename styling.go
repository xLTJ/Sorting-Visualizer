package main

import "github.com/charmbracelet/lipgloss"

var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			Padding(0, 2)

	statsStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			Width(30).
			BorderTop(false)

	helpStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			PaddingBottom(1)

	windowStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			BorderTop(false)

	windowTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#7D56F4"))

	axisStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("5"))

	blockStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("6"))
)
