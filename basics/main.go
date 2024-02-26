package main

import (
	"basics/mathFunctions"
	"basics/condFunctions"
	"basics/dataStructures"
	"fmt"
)

func main() {
	//first log
	fmt.Println("----------------------------")
	fmt.Print("helloo safeer \n")

	//variables
	fmt.Println("----------------------------")
	var name string = "safeer"
	var age int = 21

	//constants which we can't change the value
	const place string = "masterpadi"
	fmt.Printf(" %v old %v is from %v", age, name, place)

	fmt.Println("----------------------------")
	var additionResult = mathFunctions.AddIntegers(7, 7)
	var substractionResult = mathFunctions.Substract(7, 7)
	fmt.Printf("we got %v as addition result \n", additionResult)
	fmt.Printf("we got %v as substraction result \n", substractionResult)
	var isEven = condFunctions.CheckIsEven(age)
	var day = condFunctions.GetDay(7)
	fmt.Printf("then number %v is even: %v \n", age, isEven)
	fmt.Println("we got as the day for corresponding number ", day)

	// ds
	// dataStructures.DsArray()
	// dataStructures.DsSlice()
	dataStructures.DsMap()
}
