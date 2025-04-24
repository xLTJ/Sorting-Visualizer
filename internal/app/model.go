package app

import (
	"Sorting-Visualizer/internal/algorithms"
	"Sorting-Visualizer/internal/contants"
	"Sorting-Visualizer/internal/ui"
	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

const (
	stateMenu          = 0
	stateVisualization = 1
)

type Model struct {
	state     int
	help      help.Model
	list      list.Model
	sorting   bool
	canvas    canvas.Model
	sortState algorithms.SortStateInterface
}

type TickMsg time.Time

// Tick is a command sent every time we need to take a sorting step
func Tick() tea.Cmd {
	return tea.Tick(time.Millisecond*contants.SortDelay, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func NewModel() Model {
	listItems := []list.Item{
		item{title: "Bubble Sort", desc: "Bubble sort is one of the simplest sorting algorithms of all time. Its also really bad."},
		item{title: "Insertion Sort", desc: "Its insertion sort, it... inserts things :pray:"},
		item{title: "Selection Sort", desc: "Its selecting time"},
	}

	algorithmList := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	algorithmList.Title = "Sorting Algorithm Visualizer"
	algorithmList.Styles.Title = ui.TitleStyle
	algorithmList.Styles.HelpStyle = ui.HelpStyle
	algorithmList.KeyMap.CursorUp = ui.MenuKeyMap.Up
	algorithmList.KeyMap.CursorDown = ui.MenuKeyMap.Down

	visualizationCanvas := canvas.New(100, 100, canvas.WithViewWidth(0), canvas.WithViewHeight(0))

	return Model{
		state:     stateMenu,
		help:      help.New(),
		list:      algorithmList,
		sorting:   false,
		canvas:    visualizationCanvas,
		sortState: algorithms.NewSortStateBubbleSort(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
