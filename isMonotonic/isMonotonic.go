package main

import "fmt"

func isMonotonic(in []int) bool {
	less := false
	great := true

	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] && less == false {
			great = false
		} else if in[i] < in[i-1] && great == false {
			return false
		} else if in[i] < in[i-1] && great == true {
			less = true
		} else if in[i] > in[i-1] && less == true {
			return false
		}

	}

	return true
}

func main() {
	s := [][]int{
		{1, 7},
		{1, 1},
		{3, 3, 1},
		{9, 5, 1},
		{23, 5, 23},
		{23, 28, 28, 1},
		{1, 2, 3, 4, 1},
	}

	for _, val := range s {
		fmt.Println(isMonotonic(val))
	}
}
