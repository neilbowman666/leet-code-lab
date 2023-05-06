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
