package main

import "github.com/charmbracelet/bubbles/key"

type VisualizationKeyMap struct {
	Start key.Binding
	Stop  key.Binding
	Reset key.Binding
	Back  key.Binding
	Quit  key.Binding
}

type MenuKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
}

var visualizationKeyMap = VisualizationKeyMap{
	Start: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "start sorting"),
	),
	Stop: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "stop sorting"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reset array"),
	),
	Back: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "back to menu"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
}

func (v VisualizationKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{v.Start, v.Stop, v.Reset, v.Back, v.Quit}
}

func (v VisualizationKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{v.Start, v.Stop, v.Reset, v.Back, v.Quit}}
}

var menuKeyMap = MenuKeyMap{
	Up: key.NewBinding(
		key.WithKeys("w", "up"),
		key.WithHelp("↑/w", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("s", "down"),
		key.WithHelp("↓/s", "move down"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", "select"),
		key.WithHelp("enter", "select"),
	),
}
