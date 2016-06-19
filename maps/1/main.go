package main

import (
	"fmt"
)

func createPopulateMap() {

	fmt.Printf("Program started!\n")

	//create the map
	numberMap := make(map[int]int)

	//populate the map
	for i := 0; i < 10; i++ {
		numberMap[i] = i * 500
	}

	//print entire number map
	fmt.Println("Printing entire map: ", numberMap)
	fmt.Println("Length of map: ", len(numberMap))

	// checking for existence ok a key without having to iterate over the map using "comma ok" idiom
	// Refer to stack overflow http://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go?rq=1
	// val & exists are local to the if and else scopes
	if val, exists := numberMap[91]; exists {
		delete(numberMap, 91)          // delete using key
		fmt.Printf("Value: %d\n", val) // prints value using conversion specifier or placeholder
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("Value: ", val)
		fmt.Println("exists: ", exists)
	} //access to a key that doesn't exist returns zero rather than error

	//iterating over a map in an unordered fashion
	for key, val := range numberMap {
		fmt.Println(key, " - ", val)
	}

	// iterating over a map when the keys aren't required
	for _, val := range numberMap {
		fmt.Print(val, "\n")
	}

	//iterating over a map in an ordered fashion
	for keyIndex := 0; keyIndex < len(numberMap); keyIndex++ {
		fmt.Printf("Key: %d, %d\n", keyIndex, numberMap[keyIndex])
	}

}

func main() {
	createPopulateMap()
}
