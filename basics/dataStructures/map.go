package dataStructures

import (
	"fmt"
)

func DsMap() {
	// initializing without values
	fruit := make(map[string]string)
	fruit["name"] = "apple"
	fruit["from"] = "kerala"
	fruit["price"] = "170"

	// initializing with values
	fruitsWithPrices := map[string]int{
		"apple":    20,
		"grape":    90,
		"savargil": 80,
	}

	// printing the map
	for key, value := range fruitsWithPrices {
		fmt.Println(key, value)
	}

	for key, value := range fruit {
		fmt.Println(key, value)
	}
}
