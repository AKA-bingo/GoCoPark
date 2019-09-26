package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int64{44, 38, 12, 5, 44, 73, 31, 3, 66, 17, 13, 44, 51, 35}
	BubbleSort(list)
	fmt.Println(list)
}

func TestSelectSort(t *testing.T) {
	list := []int64{44, 38, 12, 5, 44, 73, 31, 3, 66, 17, 13, 44, 51, 35}
	SelectSort(list)
	fmt.Println(list)
}

func TestInsertionSort(t *testing.T) {
	list := []int64{44, 38, 12, 5, 44, 73, 31, 3, 66, 17, 13, 44, 51, 35}
	InsertionSort(list)
	fmt.Println(list)
}

func TestShellSort(t *testing.T) {
	list := []int64{44, 38, 12, 5, 44, 73, 31, 3, 66, 17, 13, 44, 51, 35}
	ShellSort(list)
	fmt.Println(list)
}

func TestQuickSort(t *testing.T) {
	list := []int64{44, 38, 12, 5, 44, 73, 31, 3, 66, 17, 13, 44, 51, 35}
	QuickSort(list)
	fmt.Println(list)
}

func TestHeapSort(t *testing.T) {
	list := []int64{44, 38, 12, 5, 44, 73, 31, 3, 66, 17, 13, 44, 51, 35}
	HeapSort(list)
	fmt.Println(list)
}
