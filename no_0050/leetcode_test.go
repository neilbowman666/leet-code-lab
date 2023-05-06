package no_0050

import (
	"fmt"
	"testing"
)

func Test0050(t *testing.T) {
	ansFunc := myPow2
	fmt.Printf(">>>>>>>>>>> leet code No.50 result sample 1: %v\n\n", ansFunc(2, 10))
	fmt.Printf(">>>>>>>>>>> leet code No.50 result sample 2: %v\n\n", ansFunc(2.1, 3))
	fmt.Printf(">>>>>>>>>>> leet code No.50 result sample 3: %v\n\n", ansFunc(2, -2))
	fmt.Printf(">>>>>>>>>>> my test: %v\n\n", ansFunc(2, 11))
}
