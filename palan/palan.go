package palan

import (
	"fmt"
)

func checkFoward(startArray []string, checkWord string) string {
	var place int = 0
	var size int = len(checkWord)

	checkLetter := startArray[place]

	if place < size {
		place++
	} else {

	}

	return checkLetter

}

func checkBackword(endArray []string, checkWord string) string {
	var size int = len(checkWord)
	var place int = size

	checkLetter := endArray[place]

	if place > 0 {
		place--
	} else {

	}

	return checkLetter

}

func CheckPalan() int {
	var word string = "hello"
	count := 0

	fmt.Print("What is the word you want to check: ")
	fmt.Scan(&word)

	size := len(word)

	arr1 := make([]string, size)

	for i := 0; i < len(word); i++ {
		arr1[i] = string(word[i])
	}

	arr2 := make([]string, size)

	for k := 0; k < size; k++ {
		arr2[k] = string(word[k])
	}

	/*//iterating foward
	for j := 0; j < size; j++ {
		start := arr1[j]
		fmt.Println(start)
	}

	//iterate backwords
	for l := len(arr2) - 1; l >= 0; l-- {
		end := arr2[l]
		fmt.Println(end)
	}*/

	//next
	/* Test out the fucntions and see if they work to what you think they will do*/

	return count
}
