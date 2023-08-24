package palan

import (
	"fmt"
)

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

	//iterating foward
	for j := 0; j < size; j++ {
		start := arr1[j]
		fmt.Println(start)
	}

	//iterating backwards
	for l := len(arr2) - 1; l >= 0; l-- {
		end := arr2[l]
		fmt.Println(end)
	}

	//next
	/* combing the both and then comparing front and back end to make sure they all match, if not not palandrome and then break out*/

	return count
}
