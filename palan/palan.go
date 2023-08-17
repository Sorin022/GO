package palan

import (
	"fmt"
	"strings"
)

func CheckPalan() string {
	var word string = "hello"
	var resultStart string = "a"
	var resultEnd string = "a"

	var count int = 0

	fmt.Print("What is the word you want to check: ")
	fmt.Scan(&word)

	for i := 0; i < len(word); i++ {
		var indexStart int = i
		charsStart := strings.Split(word, "")
		var resultStart string = charsStart[indexStart]
		for k := 0; k < 0; k-- {
			var indexEnd int = k
			charsEnd := strings.Split(word, "")
			var resultEnd string = charsEnd[indexEnd]
		}

		if resultStart == resultEnd {
			count++
		} else {
			return "Not A Palandrome"
		}

	}

	return "Palandrome"
}
