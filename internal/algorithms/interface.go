package algorithms

import (
	"Sorting-Visualizer/internal/contants"
	"Sorting-Visualizer/pkg/util"
)

var AlgorithmMap = map[string]func() SortStateInterface{
	"Bubble Sort": NewSortStateBubbleSort,
}

type SortStateInterface interface {
	GetComparisons() int
	GetSwaps() int
	IsSorted() bool
	GetActiveIndices() []int
	GetArray() []float64

	GetName() string
	Reset() SortStateInterface
	SortStep() SortStateInterface
}

type SortStateBase struct {
	comparisons   int       // total comparisons
	swaps         int       // total swaps
	sorted        bool      // whether the array has finished sorted
	activeIndices []int     // Indices currently being swapped or compared
	array         []float64 // the array being sorted
}

func (s SortStateBase) GetComparisons() int     { return s.comparisons }
func (s SortStateBase) GetSwaps() int           { return s.swaps }
func (s SortStateBase) IsSorted() bool          { return s.sorted }
func (s SortStateBase) GetActiveIndices() []int { return s.activeIndices }
func (s SortStateBase) GetArray() []float64     { return s.array }

func newBaseSortState() SortStateBase {
	return SortStateBase{
		comparisons:   0,
		swaps:         0,
		sorted:        false,
		activeIndices: []int{},
		array:         util.GenerateShuffledArray(contants.ArraySize),
	}
}
