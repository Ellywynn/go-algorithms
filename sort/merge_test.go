package sort_test

import (
	"testing"

	"github.com/ellywynn/go-algorithms/sort"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	arr := []int{2, 1, 5, 9, 0, 6, 5, 20}
	arr = sort.MergeSort(arr)
	sorted := []int{0, 1, 2, 5, 5, 6, 9, 20}
	assert.True(t, sort.CompareArrays(arr, sorted))

	arr = []int{1, 2, 3}
	arr = sort.MergeSort(arr)
	assert.True(t, sort.CompareArrays(arr, []int{1, 2, 3}))

	arr = []int{1}
	arr = sort.MergeSort(arr)
	assert.True(t, sort.CompareArrays(arr, []int{1}))

	arr = []int{}
	arr = sort.MergeSort(arr)
	assert.True(t, sort.CompareArrays(arr, []int{}))

	arr = nil
	arr = sort.MergeSort(arr)
	assert.Nil(t, arr)
}
