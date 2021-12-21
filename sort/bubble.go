package sort

// Accepts an array and changes it to sorted array in given order
func BubbleSort(array []int, order Order) {
	n := len(array)

	// Empty array or array contains one elenment
	if n < 2 || array == nil {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if order == ASC {
				if array[j] > array[j+1] {
					// Swap elements
					array[j], array[j+1] = array[j+1], array[j]
				}
			} else if order == DESC {
				if array[j] < array[j+1] {
					// Swap elements
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}
	}
}
