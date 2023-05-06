package no_0012

import (
	"fmt"
	"testing"
)

func Test0050(t *testing.T) {
	ansFunc := intToRoman
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 1: %v\n\n", ansFunc(3))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 2: %v\n\n", ansFunc(4))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 3: %v\n\n", ansFunc(9))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 4: %v\n\n", ansFunc(58))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 5: %v\n\n", ansFunc(1994))
	fmt.Printf(">>>>>>>>>>> my test: %v\n\n", ansFunc(2023))
}
