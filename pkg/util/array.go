package util

import "math/rand"

func GenerateShuffledArray(size int) []float64 {
	array := make([]float64, size)
	for i := 1; i <= size; i++ {
		array[i-1] = float64(i)
	}
	for i := range array {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
	return array
}
