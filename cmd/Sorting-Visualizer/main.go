package main

import (
	"Sorting-Visualizer/internal/app"

	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	if _, err := tea.NewProgram(app.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
