package utils

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	valSli := []int{1, 9, 9, 6, 0, 7, 2, 7, 1, 9, 9, 5, 0, 7, 3, 1}
	sortingHeap := InitSortingHeap(valSli[0])
	for _, oneVal := range valSli[1:] {
		sortingHeap.Insert(oneVal)
	}
	sortedValArr := sortingHeap.SortedTraverse()
	fmt.Println(sortedValArr)
}
