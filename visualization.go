package main

import (
	"fmt"
	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/canvas/graph"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func (m model) updateVisualization(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, visualizationKeyMap.Back):
			m.state = stateMenu
			m.sorting = false
			return m, nil

		case key.Matches(msg, visualizationKeyMap.Reset):
			m.sorting = false
			m.sortState = newSortState()
			return m, nil

		case key.Matches(msg, visualizationKeyMap.Start):
			if !m.sortState.sorted {
				m.sorting = true
				return m, tick()
			}

		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// this is the worst function of all time
func (m model) visualizationView() string {
	graphFrameW, _ := windowStyle.GetFrameSize()

	statusText := "Waiting..."
	if m.sorting {
		statusText = "Sorting..."
	} else if m.sortState.sorted {
		statusText = "Sorting complete ^~^"
	}

	// chart axis
	m.canvas.Clear()
	axisPoint := canvas.Point{X: 0, Y: m.canvas.ViewHeight}
	//graph.DrawXYAxis(&m.canvas, axisPoint, axisStyle)

	// scaling graph so it actually fits on the terminal
	columns := make([]float64, arraySize)
	for i, val := range m.sortState.array {
		columns[i] = val * float64(m.canvas.ViewHeight) / arraySize
	}

	graph.DrawColumns(&m.canvas, axisPoint, columns, blockStyle)

	graphTitle := fmt.Sprintf("╭───┤ %s - %s ├", m.algorithm, statusText)
	graphLine := strings.Repeat("─", max(0, lipgloss.Width(m.canvas.View())-lipgloss.Width(graphTitle)+graphFrameW-1))
	graphUpperBorder := lipgloss.JoinHorizontal(lipgloss.Center, graphTitle, graphLine)

	stats := statsStyle.Render(fmt.Sprintf(
		"Algorithm: %s\n\nComparisons: %d\nSwaps: %d\nArray Size: %d",
		m.algorithm,
		m.sortState.comparisons,
		m.sortState.swaps,
		len(m.sortState.array),
	))

	statsTitle := "╭───┤ Stats ├"
	statsLine := strings.Repeat("─", max(0, lipgloss.Width(stats)-lipgloss.Width(statsTitle)-1))
	statsUpperBorder := lipgloss.JoinHorizontal(lipgloss.Center, statsTitle, statsLine)

	helpView := helpStyle.Render(m.help.View(visualizationKeyMap))

	mainView := lipgloss.JoinHorizontal(
		lipgloss.Top,
		windowTitleStyle.Render(fmt.Sprintf("%s╮\n%s", graphUpperBorder, windowStyle.Render(m.canvas.View()))),
		windowTitleStyle.Render(fmt.Sprintf("%s╮\n%s", statsUpperBorder, stats)),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		mainView,
		helpView,
	)
}
