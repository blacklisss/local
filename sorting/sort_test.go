package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 2, 1, 8, 4}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 8}) {
		t.Errorf("Expected [1 2 4 5 8], but got %v", arr)
	}

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected [1 2 3 4 5 6 7 8 9 10], but got %v", arr)
	}

	arr = []int{1}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, []int{1}) {
		t.Errorf("Expected [1], but got %v", arr)
	}

	arr = []int{}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, []int{}) {
		t.Errorf("Expected [], but got %v", arr)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 2, 1, 8, 4}
	SelectionSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 8}) {
		t.Errorf("Expected [1 2 4 5 8], but got %v", arr)
	}

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected [1 2 3 4 5 6 7 8 9 10], but got %v", arr)
	}

	arr = []int{1}
	SelectionSort(arr)
	if !reflect.DeepEqual(arr, []int{1}) {
		t.Errorf("Expected [1], but got %v", arr)
	}

	arr = []int{}
	SelectionSort(arr)
	if !reflect.DeepEqual(arr, []int{}) {
		t.Errorf("Expected [], but got %v", arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 2, 1, 8, 4}
	InsertionSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 8}) {
		t.Errorf("Expected [1 2 4 5 8], but got %v", arr)
	}

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertionSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected [1 2 3 4 5 6 7 8 9 10], but got %v", arr)
	}

	arr = []int{1}
	InsertionSort(arr)
	if !reflect.DeepEqual(arr, []int{1}) {
		t.Errorf("Expected [1], but got %v", arr)
	}

	arr = []int{}
	InsertionSort(arr)
	if !reflect.DeepEqual(arr, []int{}) {
		t.Errorf("Expected [], but got %v", arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 2, 1, 8, 4}
	arr = MergeSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 8}) {
		t.Errorf("Expected [1 2 4 5 8], but got %v", arr)
	}

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	arr = MergeSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected [1 2 3 4 5 6 7 8 9 10], but got %v", arr)
	}

	arr = []int{1}
	arr = MergeSort(arr)
	if !reflect.DeepEqual(arr, []int{1}) {
		t.Errorf("Expected [1], but got %v", arr)
	}

	arr = []int{}
	arr = MergeSort(arr)
	if !reflect.DeepEqual(arr, []int{}) {
		t.Errorf("Expected [], but got %v", arr)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 2, 1, 8, 4}
	QuickSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 8}) {
		t.Errorf("Expected [1 2 4 5 8], but got %v", arr)
	}

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected [1 2 3 4 5 6 7 8 9 10], but got %v", arr)
	}

	arr = []int{1}
	QuickSort(arr)
	if !reflect.DeepEqual(arr, []int{1}) {
		t.Errorf("Expected [1], but got %v", arr)
	}

	arr = []int{}
	QuickSort(arr)
	if !reflect.DeepEqual(arr, []int{}) {
		t.Errorf("Expected [], but got %v", arr)
	}
}

func TestShellSort(t *testing.T) {
	arr := []int{5, 2, 1, 8, 4}
	ShellSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 8}) {
		t.Errorf("Expected [1 2 4 5 8], but got %v", arr)
	}

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	ShellSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected [1 2 3 4 5 6 7 8 9 10], but got %v", arr)
	}

	arr = []int{1}
	ShellSort(arr)
	if !reflect.DeepEqual(arr, []int{1}) {
		t.Errorf("Expected [1], but got %v", arr)
	}

	arr = []int{}
	ShellSort(arr)
	if !reflect.DeepEqual(arr, []int{}) {
		t.Errorf("Expected [], but got %v", arr)
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{5, 2, 1, 8, 4}
	HeapSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 8}) {
		t.Errorf("Expected [1 2 4 5 8], but got %v", arr)
	}

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	HeapSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("Expected [1 2 3 4 5 6 7 8 9 10], but got %v", arr)
	}

	arr = []int{1}
	HeapSort(arr)
	if !reflect.DeepEqual(arr, []int{1}) {
		t.Errorf("Expected [1], but got %v", arr)
	}

	arr = []int{}
	HeapSort(arr)
	if !reflect.DeepEqual(arr, []int{}) {
		t.Errorf("Expected [], but got %v", arr)
	}
}
