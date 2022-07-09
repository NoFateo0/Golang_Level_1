package inserts

func InsertionSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		var t int
		for j := i; j > 0 && (slice[j] < slice[j-1]); j-- {
			t = slice[j-1]
			slice[j-1] = slice[j]
			slice[j] = t
		}
	}
	return slice
}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
