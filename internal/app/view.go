package app

import (
	"Sorting-Visualizer/internal/contants"
	"Sorting-Visualizer/internal/ui"
	"fmt"
	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/canvas/graph"
	"github.com/charmbracelet/lipgloss"
	"slices"
	"strings"
)

func (m Model) View() string {
	switch m.state {
	case stateMenu:
		return ui.AppStyle.Render(m.list.View())
	case stateVisualization:
		return ui.AppStyle.Render(m.visualizationView())
	default:
		return "Unknown State, wtf did u do to end up here?"
	}
}

func (m Model) visualizationView() string {
	graphFrameW, _ := ui.WindowStyle.GetFrameSize()
	blockStyle := ui.InactiveBlockStyle
	swappedBlockStyle := ui.SwappedBlockStyle
	comparedBlockStyle := ui.ComparedBlockStyle

	statusText := "Waiting..."
	if m.sorting {
		statusText = "Sorting..."
		blockStyle = ui.BlockStyle
	} else if m.sortState.IsSorted() {
		statusText = "Sorting complete ^~^"
		blockStyle = ui.SortedBlockStyle
		swappedBlockStyle = ui.SortedBlockStyle
		comparedBlockStyle = ui.SortedBlockStyle
	}

	m.canvas.Clear()
	axisPoint := canvas.Point{X: 0, Y: m.canvas.ViewHeight}

	// scaling graph so it actually fits on the terminal
	columns := make([]float64, contants.ArraySize)
	for i, val := range m.sortState.GetArray() {
		// if index is active, skip to draw later
		if slices.Contains(m.sortState.GetSwappedIndices(), i) || slices.Contains(m.sortState.GetComparedIndices(), i) {
			columns[i] = 0
			continue
		}
		columns[i] = val * float64(m.canvas.ViewHeight) / contants.ArraySize
	}

	// draw compared columns
	for _, arrayIndex := range m.sortState.GetComparedIndices() {
		if arrayIndex < 0 || arrayIndex > contants.ArraySize-1 {
			continue
		}

		point := axisPoint.Add(canvas.Point{X: arrayIndex})
		scaledHeight := m.sortState.GetArray()[arrayIndex] * float64(m.canvas.ViewHeight) / contants.ArraySize
		graph.DrawColumns(&m.canvas, point, []float64{scaledHeight}, comparedBlockStyle)
	}

	// draw swapped columns
	graph.DrawColumns(&m.canvas, axisPoint, columns, blockStyle)
	for _, arrayIndex := range m.sortState.GetSwappedIndices() {
		if arrayIndex < 0 || arrayIndex > contants.ArraySize-1 {
			continue
		}

		point := axisPoint.Add(canvas.Point{X: arrayIndex})
		scaledHeight := m.sortState.GetArray()[arrayIndex] * float64(m.canvas.ViewHeight) / contants.ArraySize
		graph.DrawColumns(&m.canvas, point, []float64{scaledHeight}, swappedBlockStyle)
	}

	graphTitle := fmt.Sprintf("╭───┤ %s - %s ├", m.sortState.GetName(), statusText)
	graphLine := strings.Repeat("─", max(0, lipgloss.Width(m.canvas.View())-lipgloss.Width(graphTitle)+graphFrameW-1))
	graphUpperBorder := lipgloss.JoinHorizontal(lipgloss.Center, graphTitle, graphLine)

	stats := ui.StatsStyle.Render(fmt.Sprintf(
		"Algorithm: %s\n\nComparisons: %d\nSwaps: %d\nArray Size: %d",
		m.sortState.GetName(),
		m.sortState.GetComparisons(),
		m.sortState.GetSwaps(),
		len(m.sortState.GetArray()),
	))

	statsTitle := "╭───┤ Stats ├"
	statsLine := strings.Repeat("─", max(0, lipgloss.Width(stats)-lipgloss.Width(statsTitle)-1))
	statsUpperBorder := lipgloss.JoinHorizontal(lipgloss.Center, statsTitle, statsLine)

	helpView := ui.HelpStyle.Render(m.help.View(ui.VisualizationKeyMap))

	mainView := lipgloss.JoinHorizontal(
		lipgloss.Top,
		ui.WindowTitleStyle.Render(fmt.Sprintf("%s╮\n%s", graphUpperBorder, ui.WindowStyle.Render(m.canvas.View()))),
		ui.WindowTitleStyle.Render(fmt.Sprintf("%s╮\n%s", statsUpperBorder, stats)),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		mainView,
		helpView,
	)
}
