package numbers

import "fmt"

func AddTwoNums() int {
	var x int = 0
	var y int = 0

	fmt.Print("What is your first num: ")
	fmt.Scan(&x)
	fmt.Print("What is your second num: ")
	fmt.Scan(&y)

	return x + y
}
