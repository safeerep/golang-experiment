package dataStructures

import (
	"fmt"
)

func DsSlice () {
	var numbers []int
	numbers = append(numbers, 12, 23, 34, 45, 56, 67)
	fmt.Println(numbers)
	
	// removing first element
	numbers = append(numbers[:0], numbers[1:]...)
	fmt.Println(numbers)

	// removing last element
	numbers = numbers[:len(numbers) - 1]
	fmt.Println(numbers)

	// to remove a specific element
	var newSlice []int
	for _, number := range numbers {
		if number != 34 {
			newSlice = append(newSlice, number)
		}
	}
	numbers = newSlice
	fmt.Println(numbers)
}