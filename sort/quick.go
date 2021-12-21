package sort

func QuickSort(array []int) []int {
	return quick(array, 0, len(array)-1)
}

func partition(array []int, low, high int) ([]int, int) {
	pivot := array[high] // The last element of the part

	// Move all values which are less than pivot to the left of pivot
	i := low
	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}

	array[i], array[high] = array[high], array[i]
	return array, i
}

func quick(array []int, low, high int) []int {
	if low < high {
		var p int
		array, p = partition(array, low, high)
		array = quick(array, low, p-1)
		array = quick(array, p+1, high)
	}
	return array
}
