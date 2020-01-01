package mysort

func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func InsertionSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				swap(values, i, j)
			}
		}
	}
}

func swap(values []int, i int, j int) {
	tmp := values[i]
	values[i] = values[j]
	values[j] = tmp
}
