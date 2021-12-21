package sort_test

import (
	"testing"

	"github.com/ellywynn/go-algorithms/sort"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 1, 10, 0, 4, 2, 9}
	arr = sort.QuickSort(arr)
	assert.True(t, sort.CompareArrays(arr, []int{0, 1, 2, 2, 4, 9, 10}))

	arr = []int{1}
	assert.True(t, sort.CompareArrays(arr, []int{1}))

	arr = []int{}
	assert.True(t, sort.CompareArrays(arr, []int{}))

	arr = nil
	assert.Nil(t, arr)
}
