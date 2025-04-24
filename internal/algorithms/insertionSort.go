package algorithms

type SortStateInsertionSort struct {
	SortStateBase
	i   int
	j   int
	key int
}

func NewSortStateInsertionSort() SortStateInterface {
	return SortStateInsertionSort{
		SortStateBase: newBaseSortState(),
		i:             0,
		j:             0,
		key:           0,
	}
}

func (s SortStateInsertionSort) GetName() string {
	return "Insertion Sort"
}

func (s SortStateInsertionSort) Reset() SortStateInterface {
	return NewSortStateInsertionSort()
}

func (s SortStateInsertionSort) SortStep() SortStateInterface {
	// grabs current state
	arr := state.array
	n := len(arr)
	i, j := state.i, state.j

	key := arr[i]
}
