package main

import (
	"fmt"
)

func main() {

	//declare and initialise a map
	myMap := make(map[string]int)

	//inserting key value pairs into the map
	myMap["apple"] = 1
	myMap["banana"] = 2
	myMap["orange"] = 3

	//accessing values using keys
	appleValue := myMap["apple"]
	bananaValue := myMap["banana"]
	fmt.Println("Value of apple:", appleValue)
	fmt.Println("Value of banana:", bananaValue)

	//updating values
	myMap["apple"] = 5
	fmt.Println("updated value of apple:", myMap["apple"])

	//deleting a key value
	delete(myMap, "orange")
	fmt.Println("after deleting orange:", myMap)

	//checking existence of a key
	value, exists := myMap["banana"]
	if exists {
		fmt.Println("Value of banana:", value)
	} else {
		fmt.Println("Banana not found in the map")
	}

	//iterating over the map
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	//length of the map
	fmt.Println("LEngth of the map:", len(myMap))
}
