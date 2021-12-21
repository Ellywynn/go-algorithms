package sort

// Sorts an array(mutable)
func SelectionSort(array []int, order Order) {
	n := len(array)

	// If array contains one or zero elements or array is nil, do nothing
	if n < 2 || array == nil {
		return
	}

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if order == ASC {
				if array[j] < array[minIndex] {
					minIndex = j
				}
			} else {
				if array[j] > array[minIndex] {
					minIndex = j
				}
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}
