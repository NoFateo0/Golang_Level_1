package primen

func PrimeN(n int) (result []int) {
	for i := 2; i <= n; i++ {
		flag := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
			}
		}
		if flag == true {
			continue
		} else {
			result = append(result, i)
		}
	}
	return
}
