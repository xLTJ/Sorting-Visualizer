package algorithms

type SortStateInsertionSort struct {
	SortStateBase
	i     int
	j     int
	key   float64
	phase SortPhase
}

func NewSortStateInsertionSort() SortStateInterface {
	return SortStateInsertionSort{
		SortStateBase: newBaseSortState(),
		i:             1,
		j:             0,
		key:           0,
		phase:         outerLoopStart,
	}
}

func (s SortStateInsertionSort) GetName() string {
	return "Insertion Sort"
}

func (s SortStateInsertionSort) Reset() SortStateInterface {
	return NewSortStateInsertionSort()
}

func (s SortStateInsertionSort) SortStep() SortStateInterface {
	if s.sorted {
		return s
	}

	s.comparedIndices = []int{}
	s.swappedIndices = []int{}

	// grabs current state
	arr := s.array
	n := len(arr)
	i, j, key, phase := s.i, s.j, s.key, s.phase

	switch phase {
	case outerLoopStart:
		key = arr[i]
		j = i - 1

		s.comparedIndices = []int{i, j}
		s.comparisons++
		if arr[j] > key {
			phase = innerLoop
		} else {
			phase = outerLoopEnd
		}

	case innerLoop:
		arr[j+1] = arr[j]
		s.swappedIndices = []int{j + 1, j}
		s.swaps++
		j--
		if !(j >= 0 && arr[j] > key) {
			phase = outerLoopEnd
		}
		s.comparedIndices = []int{j, i}
		s.comparisons++

	case outerLoopEnd:
		arr[j+1] = key
		i++
		phase = outerLoopStart
		s.swappedIndices = []int{j + 1}
		s.swaps++
	}

	if i >= n {
		s.sorted = true
	}

	s.i, s.j, s.key, s.phase = i, j, key, phase
	return s
}
