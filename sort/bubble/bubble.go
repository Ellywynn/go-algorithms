package bubble

type Order int

const (
	ASC Order = iota
	DESC
)

// Accepts an array and changes it to sorted array in given order
func BubbleSort(array []int, order Order) {
	n := len(array)

	// Empty array or array contains one elenment
	if n < 2 || array == nil {
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if order == ASC {
				if array[i] < array[j] {
					// Swap elements
					array[i], array[j] = array[j], array[i]
				}
			} else if order == DESC {
				if array[i] > array[j] {
					// Swap elements
					array[i], array[j] = array[j], array[i]
				}
			}
		}
	}
}
