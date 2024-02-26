package dataStructures

import (
	"fmt"
)

func DsArray () {
	// first array without values specifying on declaration time;
	var intArray [10]int
	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3
	fmt.Println(intArray)

	// values specifying on array declaration time
	stringArray := [3]string{"first", "second", "third"}
	fmt.Println(stringArray)
	lastValueOfStringArray := stringArray[len(stringArray) - 1]
	fmt.Println("last value is", lastValueOfStringArray)

	// for loop
	for i := 0; i < len(stringArray); i++ {
		fmt.Println(stringArray[i])
	}

	// for loop concise model range based;
	// the first underscore' position is where index value comes;
	for _, num := range intArray {
		fmt.Println(num * 2)
	}

	// array slicing
	fmt.Println(intArray[1:2])
}