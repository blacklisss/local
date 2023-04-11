package sorting

// QuickSort сортирует переданный массив методом быстрой сортировки
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	left, right := 1, len(arr)-1
	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++
		}
		for right >= left && arr[right] >= pivot {
			right--
		}
		if right >= left {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[0], arr[right] = arr[right], arr[0]
	QuickSort(arr[:right])
	QuickSort(arr[right+1:])
	return arr
}
