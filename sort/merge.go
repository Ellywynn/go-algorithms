package sort

// Sorts an array(immutable)
func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}

	left := MergeSort(array[:n/2])
	right := MergeSort(array[n/2:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	sorted := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		sorted = append(sorted, left[i])
	}
	for ; j < len(right); j++ {
		sorted = append(sorted, right[j])
	}
	return sorted
}
