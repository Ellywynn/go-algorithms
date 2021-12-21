package sort_test

import (
	"testing"

	"github.com/ellywynn/go-algorithms/sort"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 1, 4, 2, 8}
	sorted := []int{1, 2, 4, 5, 8}
	sort.BubbleSort(arr, sort.ASC)
	assert.True(t, sort.CompareArrays(arr, sorted))

	sorted = []int{8, 5, 4, 2, 1}
	sort.BubbleSort(arr, sort.DESC)
	assert.True(t, sort.CompareArrays(arr, sorted))

	arr = []int{}
	sort.BubbleSort(arr, sort.ASC)
	assert.True(t, sort.CompareArrays(arr, []int{}))

	arr = nil
	sort.BubbleSort(arr, sort.ASC)
	assert.Nil(t, arr)
}
