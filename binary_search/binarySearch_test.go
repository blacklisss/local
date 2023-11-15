package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}

	// Проверка существующего элемента
	if binarySearch(arr, 5) != 2 {
		t.Error("Ожидалось, что 5 будет найден в индексе 2")
	}

	// Проверка отсутствующего элемента
	if binarySearch(arr, 4) != -1 {
		t.Error("Ожидалось, что 4 не будет найден")
	}

	// Проверка первого элемента
	if binarySearch(arr, 1) != 0 {
		t.Error("Ожидалось, что первый элемент будет найден в индексе 0")
	}

	// Проверка последнего элемента
	if binarySearch(arr, 9) != 4 {
		t.Error("Ожидалось, что последний элемент будет найден в индексе 4")
	}

	// Проверка пустого массива
	if binarySearch([]int{}, 1) != -1 {
		t.Error("Ожидалось, что в пустом массиве ничего не будет найдено")
	}
}

func TestBinarySearchAsm(t *testing.T) {
	arr := []int64{1, 3, 5, 7, 9}

	res := BinarySearch(arr, int64(len(arr)), 0)
	fmt.Println(res)
	/*// Проверка существующего элемента
	if BinarySearch(&arr, int64(len(arr)), 5) != 2 {
		t.Error("Ожидалось, что 5 будет найден в индексе 2")
	}*/

	/*// Проверка отсутствующего элемента
	if BinarySearch(arr, int64(len(arr)), 4) != -1 {
		t.Error("Ожидалось, что 4 не будет найден")
	}

	// Проверка первого элемента
	if BinarySearch(arr, int64(len(arr)), 1) != 0 {
		t.Error("Ожидалось, что первый элемент будет найден в индексе 0")
	}

	// Проверка последнего элемента
	if BinarySearch(arr, int64(len(arr)), 9) != 4 {
		t.Error("Ожидалось, что последний элемент будет найден в индексе 4")
	}

	// Проверка пустого массива
	if BinarySearch([]int64{}, int64(len(arr)), 1) != -1 {
		t.Error("Ожидалось, что в пустом массиве ничего не будет найдено")
	}*/
}

func BenchmarkBinarySearchGo(b *testing.B) {
	arr := []int{1, 3, 5, 7, 9}
	for n := 0; n < b.N; n++ {
		_ = binarySearch(arr, 5)
	}
}

//go:noescape
func BinarySearch(array []int64, length, target int64) int64

func BenchmarkBinarySearchAsm(b *testing.B) {
	array := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := int64(5)

	for n := 0; n < b.N; n++ {
		_ = BinarySearch(array, int64(len(array)), target)
	}
}
