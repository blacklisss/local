package binary_search

func binarySearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
