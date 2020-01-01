package mysort

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func quickSort(values []int, left int, right int) {
	if right-left <= 1 {
		return
	}
	//partition
	pvalue := values[right]
	p := left
	for i := left; i < right; i++ {
		if values[i] < pvalue {
			if p < i {
				values[i], values[p] = values[p], values[i]
			}
			p++ //索引小于p的元素都小于 pvalue
		}
	}
	values[p], values[right] = values[right], values[p]
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}
