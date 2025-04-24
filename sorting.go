package main

func bubbleSortStep(state SortState) SortState {
	array := state.array
	n := len(array)
	i, j := state.i, state.j

	if state.sorted {
		return state
	}

	state.activeIndices = []int{j, j + 1}

	if array[j] > array[j+1] {
		array[j], array[j+1] = array[j+1], array[j]
		state.swaps++
	}
	state.comparisons++

	j++
	if j >= n-i-1 {
		i++
		j = 0
	}

	if i >= n-1 {
		state.sorted = true
	}

	state.i = i
	state.j = j

	return state
}
