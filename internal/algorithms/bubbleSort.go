package algorithms

type SortStateBubbleSort struct {
	SortStateBase
	i int
	j int
}

func NewSortStateBubbleSort() SortStateInterface {
	return SortStateBubbleSort{
		i:             0,
		j:             0,
		SortStateBase: newBaseSortState(),
	}
}

func (s SortStateBubbleSort) Reset() SortStateInterface {
	return NewSortStateBubbleSort()
}

func (s SortStateBubbleSort) GetName() string {
	return "Bubble Sort"
}

func (s SortStateBubbleSort) SortStep() SortStateInterface {
	if s.sorted {
		return s
	}

	s.comparedIndices = []int{}
	s.swappedIndices = []int{}

	arr := s.array
	n := len(arr)
	i, j := s.i, s.j

	s.comparedIndices = []int{j, j + 1}

	if arr[j] > arr[j+1] {
		arr[j], arr[j+1] = arr[j+1], arr[j]
		s.swappedIndices = []int{j, j + 1}
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
