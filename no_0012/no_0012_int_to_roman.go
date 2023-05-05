package no_0012

import "fmt"

func digitToRoman(digitNum int, basicRoman string, fiveRoman string, tenRoman string) string {

	romanDigit := ""

	if digitNum <= 3 && digitNum > 0 {
		for i := 0; i < digitNum-0; i++ {
			romanDigit += basicRoman
		}
	} else if digitNum == 4 {
		romanDigit = basicRoman + fiveRoman
	} else if digitNum < 9 && digitNum >= 5 {
		romanDigit = fiveRoman
		for i := 0; i < digitNum-5; i++ {
			romanDigit += basicRoman
		}
	} else if digitNum == 9 {
		romanDigit = basicRoman + tenRoman
	}

	return romanDigit
}

func intToRoman(num int) string {
	var (
		digitK    = (num % 10000) / 1000
		digitH    = (num % 1000) / 100
		digitD    = (num % 100) / 10
		digitReal = (num % 10) / 1
	)
	var (
		romanK    = ""
		romanH    = ""
		romanD    = ""
		romanReal = ""
	)

	for i := 0; i < digitK; i++ {
		romanK += "M"
	}
	romanH = digitToRoman(digitH, "C", "D", "M")
	romanD = digitToRoman(digitD, "X", "L", "C")
	romanReal = digitToRoman(digitReal, "I", "V", "X")

	return fmt.Sprintf("%v%v%v%v", romanK, romanH, romanD, romanReal)
}

func Test() {
	ansFunc := intToRoman
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 1: %v\n\n", ansFunc(3))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 2: %v\n\n", ansFunc(4))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 3: %v\n\n", ansFunc(9))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 4: %v\n\n", ansFunc(58))
	fmt.Printf(">>>>>>>>>>> leet code No.12 result sample 5: %v\n\n", ansFunc(1994))
	fmt.Printf(">>>>>>>>>>> my test: %v\n\n", ansFunc(2023))
}
