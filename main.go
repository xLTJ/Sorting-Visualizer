package main

import (
	"fmt"
	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"os"
	"time"
)

const (
	stateMenu          = 0
	stateVisualization = 1
)

const (
	arraySize = 100
	sortDelay = 1
)

type Algorithm string

const (
	BubbleSort    Algorithm = "Bubble Sort"
	SelectionSort Algorithm = "Selection Sort"
)

// states -------------------------------------------------------------------------------------------------------

type model struct {
	state     int
	help      help.Model
	list      list.Model
	algorithm Algorithm // algorithm currently used
	sorting   bool
	canvas    canvas.Model
	sortState SortState
}

func newSortState() SortState {
	return SortState{
		i:             0,
		j:             0,
		comparisons:   0,
		swaps:         0,
		sorted:        false,
		activeIndices: []int{},
		array:         generateShuffledArray(arraySize),
	}
}

// SortState probably has to be expanded for more advanced algorithms but i can just add that later
type SortState struct {
	i             int       // out loop index
	j             int       // inner loop index
	comparisons   int       // total comparisons
	swaps         int       // total swaps
	sorted        bool      // whether the array has finished sorted
	activeIndices []int     // Indices currently being swapped or compared
	array         []float64 // the array being sorted
}

// setup ---------------------------------------------------------------------------------------------------

type tickMsg time.Time

// tick is a command sent every time we need to take a sorting step
func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*sortDelay, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func newModel() model {
	listItems := []list.Item{
		item{title: "Bubble Sort", desc: "Bubble sort is one of the simplest sorting algorithms of all time. Its also really bad."},
		item{title: "Insetion Sort", desc: "Its insertion sort, it... inserts things :pray:"},
	}

	algorithmList := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	algorithmList.Title = "Sorting Algorithm Visualizer"
	algorithmList.Styles.Title = titleStyle
	algorithmList.Styles.HelpStyle = helpStyle
	algorithmList.KeyMap.CursorUp = menuKeyMap.Up
	algorithmList.KeyMap.CursorDown = menuKeyMap.Down

	visualizationCanvas := canvas.New(100, 100, canvas.WithViewWidth(0), canvas.WithViewHeight(0))

	return model{
		state:     stateMenu,
		help:      help.New(),
		list:      algorithmList,
		algorithm: BubbleSort,
		sorting:   false,
		canvas:    visualizationCanvas,
		sortState: newSortState(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// updates -----------------------------------------------------------------------------------------------

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// general updates
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		appWidth, appHeight := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-appWidth, msg.Height-appHeight)

		graphWidth, graphHeight := windowStyle.GetFrameSize()
		helpHeight := helpStyle.GetHeight()
		statsWidth := statsStyle.GetWidth()

		subFromWidth := graphWidth + statsWidth
		subFromHeight := graphHeight + helpHeight + 2 // tfw magic number

		m.canvas.ViewWidth, m.canvas.ViewHeight = msg.Width-appWidth-subFromWidth, msg.Height-appHeight-subFromHeight

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, visualizationKeyMap.Quit):
			return m, tea.Quit
		}

	case tickMsg:
		if !m.sorting && m.sortState.sorted {
			break
		}

		switch m.algorithm {
		case BubbleSort:
			newState := bubbleSortStep(m.sortState)
			m.sortState = newState

			if !m.sortState.sorted {
				return m, tick()
			}
			m.sorting = false
			m.sortState.activeIndices = []int{}
		}
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

func (m model) updateMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, menuKeyMap.Select):
			if i, ok := m.list.SelectedItem().(item); ok {
				m.algorithm = Algorithm(i.Title())
				m.state = stateVisualization
				m.sortState = newSortState()
				return m, nil
			}

		case key.Matches(msg, menuKeyMap.Up):
			m.list.CursorUp()
			return m, nil

		case key.Matches(msg, menuKeyMap.Down):
			m.list.CursorDown()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// view ------------------------------------------------------------------------------------------------------

func (m model) View() string {
	switch m.state {
	case stateMenu:
		return appStyle.Render(m.list.View())
	case stateVisualization:
		return appStyle.Render(m.visualizationView())
	default:
		return "Unknown State, wtf did u do to end up here?"
	}
}

// helpers (might move these into a separate file if theres enough) ----------------------------------------------

func generateShuffledArray(size int) []float64 {
	array := make([]float64, size)
	for i := 0; i < size; i++ {
		array[i] = float64(i)
	}
	for i := range array {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
	return array
}

// main -----------------------------------------------------------------------------------------------------

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
