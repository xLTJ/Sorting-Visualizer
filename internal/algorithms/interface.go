package algorithms

import (
	"Sorting-Visualizer/internal/contants"
	"Sorting-Visualizer/pkg/util"
)

var AlgorithmMap = map[string]func() SortStateInterface{
	"Bubble Sort":    NewSortStateBubbleSort,
	"Insertion Sort": NewSortStateInsertionSort,
	"Selection Sort": NewSortStateSelectionSort,
}

type SortPhase int

const (
	outerLoopStart SortPhase = 0
	innerLoop      SortPhase = 1
	outerLoopEnd   SortPhase = 2
)

type SortStateInterface interface {
	GetComparisons() int
	GetSwaps() int
	IsSorted() bool
	GetSwappedIndices() []int
	GetComparedIndices() []int
	GetArray() []float64

	GetName() string
	Reset() SortStateInterface
	SortStep() SortStateInterface
}

type SortStateBase struct {
	comparisons     int       // total comparisons
	swaps           int       // total swaps
	sorted          bool      // whether the array has finished sorted
	swappedIndices  []int     // Indices currently being swapped
	comparedIndices []int     // indices currently being compared
	array           []float64 // the array being sorted
}

func (s SortStateBase) GetComparisons() int       { return s.comparisons }
func (s SortStateBase) GetSwaps() int             { return s.swaps }
func (s SortStateBase) IsSorted() bool            { return s.sorted }
func (s SortStateBase) GetSwappedIndices() []int  { return s.swappedIndices }
func (s SortStateBase) GetComparedIndices() []int { return s.comparedIndices }
func (s SortStateBase) GetArray() []float64       { return s.array }

func newBaseSortState() SortStateBase {
	return SortStateBase{
		comparisons:    0,
		swaps:          0,
		sorted:         false,
		swappedIndices: []int{},
		array:          util.GenerateShuffledArray(contants.ArraySize),
	}
}
