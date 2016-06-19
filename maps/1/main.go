package main

import (
	"fmt"
)

func createPopulateMap() {

	fmt.Printf("Program started!\n")

	//create the map
	numberMap := make(map[int]int)

	//populate the map
	for i := 0; i < 50; i++ {
		numberMap[i] = i * 500
	}

	//print entire number map
	fmt.Println(numberMap)
	fmt.Println("Length of map: ", len(numberMap))

	//checking for existence using "comma ok" idiom
	if val, exists := numberMap[91]; exists {
		delete(numberMap, 91)
		//prints value using conversion specifier or placeholder
		fmt.Printf("Value: %d\n", val)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("Value: ", val)
		fmt.Println("exists: ", exists)
	} //access to a key that doesn't exist returns zero rather than error

	//iterating over a map in an unordered fashion
	for key, val := range numberMap {
		fmt.Println(key, " - ", val)
	}

	//iterating over a map in an ordered fashion
	for keyIndex := 0; keyIndex < len(numberMap); keyIndex++ {
		fmt.Printf("Key: %d, %d\n", keyIndex, numberMap[keyIndex])

	}

	// Maps - shorthand notation for creating maps

	// var myMap map[string]string
	// myMap =  map[string]string{}

	myGreeting := map[string]string{
		"Tim":    "Good morning",
		"Josh":   "Good afternoon",
		"Julian": "Good night",
	}
	// prints entire map
	fmt.Println(myGreeting)

	//map and structs
}

func main() {
	createPopulateMap()
}
