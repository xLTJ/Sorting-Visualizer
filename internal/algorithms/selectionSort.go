package algorithms

type SortStateSelectionSort struct {
	SortStateBase
	i      int
	j      int
	minIdx int
	phase  SortPhase
}

func NewSortStateSelectionSort() SortStateInterface {
	return SortStateSelectionSort{
		SortStateBase: newBaseSortState(),
		i:             0,
		j:             0,
		minIdx:        0,
		phase:         outerLoopStart,
	}
}

func (s SortStateSelectionSort) GetName() string {
	return "Selection Sort"
}

func (s SortStateSelectionSort) Reset() SortStateInterface {
	return NewSortStateSelectionSort()
}

func (s SortStateSelectionSort) SortStep() SortStateInterface {
	if s.sorted {
		return s
	}

	s.comparedIndices = []int{}
	s.swappedIndices = []int{}

	// current state
	arr := s.array
	n := len(arr)
	i, j, minIdx, phase := s.i, s.j, s.minIdx, s.phase

	switch phase {
	case outerLoopStart:
		minIdx = i
		j = i + 1
		phase = innerLoop

	case innerLoop:
		if j >= n {
			phase = outerLoopEnd
			break
		}

		s.comparisons++
		s.comparedIndices = []int{j, minIdx}
		if arr[j] < arr[minIdx] {
			minIdx = j
		}
		j++

	case outerLoopEnd:
		s.swaps++
		s.swappedIndices = []int{minIdx, i}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
		i++
		phase = outerLoopStart
	}

	if i >= n {
		s.sorted = true
	}

	s.i, s.j, s.minIdx, s.phase = i, j, minIdx, phase
	return s
}
