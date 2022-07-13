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

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
