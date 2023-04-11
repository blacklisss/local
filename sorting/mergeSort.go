package sorting

// MergeSort (Сортировка слиянием) сортирует переданный массив методом слияния
func MergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	arr := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			arr[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			arr[k] = left[i]
			i++
		} else if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return arr
}
