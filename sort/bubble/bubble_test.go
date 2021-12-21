package bubble_test

import (
	"testing"

	"github.com/ellywynn/go-algorithms/sort/bubble"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 1, 4, 2, 8}
	sorted := []int{1, 2, 4, 5, 8}
	bubble.BubbleSort(arr, bubble.ASC)
	assert.True(t, compare(arr, sorted))

	sorted = []int{8, 5, 4, 2, 1}
	bubble.BubbleSort(arr, bubble.DESC)
	assert.True(t, compare(arr, sorted))

	arr = []int{}
	bubble.BubbleSort(arr, bubble.ASC)
	assert.True(t, compare(arr, []int{}))

	arr = nil
	bubble.BubbleSort(arr, bubble.ASC)
	assert.Nil(t, arr)
}

// Compares two arrays(the lenght must be the same)
func compare(arr1 []int, arr2 []int) bool {
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}
