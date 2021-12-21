package sort_test

import (
	"testing"

	"github.com/ellywynn/go-algorithms/sort"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{64, 25, 12, 22, 11}
	sort.SelectionSort(arr, sort.ASC)
	assert.True(t, sort.CompareArrays(arr, []int{11, 12, 22, 25, 64}))

	sort.SelectionSort(arr, sort.DESC)
	assert.True(t, sort.CompareArrays(arr, []int{64, 25, 22, 12, 11}))

	arr = []int{1}
	sort.SelectionSort(arr, sort.ASC)
	assert.Equal(t, 1, arr[0])

	arr = nil
	sort.SelectionSort(arr, sort.ASC)
	assert.Nil(t, arr)
}
