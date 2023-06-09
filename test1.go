package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums1 := []int{9, 2, 3}
	nums2 := []int{8, 1}

	res := add2Nums(nums1, nums2)

	fmt.Println("res = ", res)

	nums := []int{1, 0, 2, 3, 1, 2, 1, 1, 2, 2, 3}

	res = kMaxCoutn(nums, 2)

	fmt.Println(res)
}

func add2Nums(nums1 []int, nums2 []int) []int {
	i1 := len(nums1) - 1
	i2 := len(nums2) - 1

	res := make([]int, 0)

	carry := 0

	for i1 >= 0 || i2 >= 0 || carry != 0 {
		if i1 >= 0 {
			carry += nums1[i1]
			i1--
		}
		if i2 >= 0 {
			carry += nums2[i2]
			i2--
		}

		res = append(res, carry%10)
		carry /= 10
	}

	reverse(res)

	return res
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		j := l - i - 1
		nums[i], nums[j] = nums[j], nums[i]
	}
}

type pair struct {
	val   int
	count int
}

type heapPairs struct {
	pairs []pair
}

func NewHeapPairs(capacity int) *heapPairs {
	return &heapPairs{
		pairs: make([]pair, 0, capacity),
	}
}

func (h *heapPairs) Len() int           { return len(h.pairs) }
func (h *heapPairs) Less(i, j int) bool { return h.pairs[i].count > h.pairs[j].count }
func (h *heapPairs) Swap(i, j int)      { h.pairs[i], h.pairs[j] = h.pairs[j], h.pairs[i] }
func (h *heapPairs) Push(v any)         { h.pairs = append(h.pairs, v.(pair)) }
func (h *heapPairs) Pop() any {
	l := h.Len()
	x := h.pairs[l-1]
	h.pairs = h.pairs[:l-1]
	return x
}

func kMaxCoutn(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	h := NewHeapPairs(len(m))

	for k, v := range m {
		heap.Push(h, pair{k, v})
	}

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(pair).val
	}

	return res
}
