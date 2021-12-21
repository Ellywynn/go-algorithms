package sort

type Order int

const (
	ASC Order = iota
	DESC
)

// Compares two arrays
func CompareArrays(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}
