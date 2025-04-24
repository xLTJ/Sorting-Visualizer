package main

type SortStateBubbleSort struct {
	SortStateBase
	i int
	j int
}

func newSortStateBubbleSort() SortStateInterface {
	return SortStateBubbleSort{
		i:             0,
		j:             0,
		SortStateBase: newBaseSortState(),
	}
}

func (s SortStateBubbleSort) reset() SortStateInterface {
	return SortStateBubbleSort{
		i:             0,
		j:             0,
		SortStateBase: newBaseSortState(),
	}
}

func (s SortStateBubbleSort) getName() string {
	return "Bubble Sort"
}

func (s SortStateBubbleSort) sortStep() SortStateInterface {
	arr := s.array
	n := len(arr)
	i, j := s.i, s.j

	if s.sorted {
		return s
	}

	s.activeIndices = []int{j, j + 1}

	if arr[j] > arr[j+1] {
		arr[j], arr[j+1] = arr[j+1], arr[j]
		s.swaps++
	}
	s.comparisons++

	j++
	if j >= n-i-1 {
		i++
		j = 0
	}

	if i >= n-1 {
		s.sorted = true
	}

	s.i, s.j = i, j
	return s
}

type SortStateInsertionSort struct {
	SortStateBase
	i int
	j int
}

//
//func insertionSortStep(state SortState) SortState {
//	// grabs current state
//	arr := state.array
//	n := len(arr)
//	i, j := state.i, state.j
//
//	key := arr[i]
//}
