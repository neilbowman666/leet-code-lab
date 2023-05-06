package utils

import (
	"log"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{7, 4, 5, 1, 0, 9, 2, 3, 4}
	BubbleSort(&arr, false)
	log.Printf("sorted: %v\n", arr)
}
