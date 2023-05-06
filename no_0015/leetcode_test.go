package no_0015

import (
	"fmt"
	"testing"
)

func Test0050(t *testing.T) {
	ansFunc := threeSum
	fmt.Printf(">>>>>>>>>>> leet code No.50 result sample 1: %v\n\n", ansFunc([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf(">>>>>>>>>>> leet code No.50 result sample 2: %v\n\n", ansFunc([]int{0, 1, 1}))
	fmt.Printf(">>>>>>>>>>> leet code No.50 result sample 3: %v\n\n", ansFunc([]int{0, 0, 0}))
	fmt.Printf(">>>>>>>>>>> my test: %v\n\n", ansFunc([]int{-7, -4, -1, 0, 1, 2, 3, 4, 5, 6}))
}
