package main

import (
	"fmt"
)

func createPopulateMap() {

	fmt.Printf("Program started!\n")

	//create the map
	//numberMap := make(map[int]int)

	var numberMap map[int]int
	numberMap = map[int]int{}

	//populate the map
	for i := 0; i < 50; i++ {
		numberMap[i] = i * 500
	}

	//print entire number map
	fmt.Println(numberMap)
	fmt.Println("Length of map: ", len(numberMap))

	val, ok := numberMap[93]

	fmt.Println(val, ok)

	if ok {
		fmt.Println("...running..ok as true")
	} else {
		fmt.Printf("...running..ok as %t %d ", ok, val)
	}

}

func main() {
	createPopulateMap()
}
