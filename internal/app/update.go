package app

import (
	"Sorting-Visualizer/internal/algorithms"
	"Sorting-Visualizer/internal/ui"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// general updates
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		appWidth, appHeight := ui.AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-appWidth, msg.Height-appHeight)

		graphWidth, graphHeight := ui.WindowStyle.GetFrameSize()
		helpHeight := ui.HelpStyle.GetHeight()
		statsWidth := ui.StatsStyle.GetWidth()

		subFromWidth := graphWidth + statsWidth
		subFromHeight := graphHeight + helpHeight + 2 // tfw magic number

		m.canvas.ViewWidth, m.canvas.ViewHeight = msg.Width-appWidth-subFromWidth, msg.Height-appHeight-subFromHeight

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, ui.VisualizationKeyMap.Quit):
			return m, tea.Quit
		}

	case TickMsg:
		if !m.sorting || m.sortState.IsSorted() {
			break
		}

		newState := m.sortState.SortStep()
		m.sortState = newState

		if !m.sortState.IsSorted() {
			return m, Tick()
		}
		m.sorting = false
	}

	// specific state updates
	switch m.state {
	case stateMenu:
		return m.updateMenu(msg)
	case stateVisualization:
		return m.updateVisualization(msg)
	}

	return m, nil
}

func (m Model) updateMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, ui.MenuKeyMap.Select):
			if i, ok := m.list.SelectedItem().(item); ok {
				m.state = stateVisualization
				m.sortState = algorithms.AlgorithmMap[i.title]()
				return m, nil
			}

		case key.Matches(msg, ui.MenuKeyMap.Up):
			m.list.CursorUp()
			return m, nil

		case key.Matches(msg, ui.MenuKeyMap.Down):
			m.list.CursorDown()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) updateVisualization(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, ui.VisualizationKeyMap.Back):
			m.state = stateMenu
			m.sorting = false
			return m, nil

		case key.Matches(msg, ui.VisualizationKeyMap.Reset):
			m.sorting = false
			m.sortState = m.sortState.Reset()
			return m, nil

		case key.Matches(msg, ui.VisualizationKeyMap.Start):
			if !m.sortState.IsSorted() {
				m.sorting = true
				return m, Tick()
			}

		case key.Matches(msg, ui.VisualizationKeyMap.Stop):
			m.sorting = false
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
