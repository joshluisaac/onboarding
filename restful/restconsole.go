package main

import (
	"encoding/json"
	"fmt"
	//"encoding/xml"
)

type Payload struct {
	ParentMap Data
}

type Data struct {
	Fruit     Fruits
	Vegetable Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func getJsonPayloadData() ([]byte, error) {
	fruits := map[string]int{"Apples": 10, "Oranges": 50, "Mangoes": 9}
	vegetables := map[string]int{
		"Carrot":     19,
		"Brocoli":    0,
		"Beansprout": 13,
		"Spinach":    9,
		"Cabbage":    17,
		"Celery":     4,
		"Cucumbers":  7,
		"Tomatoes":   18,
		"Onion":      8}
	d := Data{fruits, vegetables}
	fmt.Println(d)
	p := Payload{d}
	return json.MarshalIndent(p, "", " ")
	//return xml.MarshalIndent(p, "", " ")
}

func servePayloadData() {
	response, err := getJsonPayloadData()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println(err)
		fmt.Println("Printing slice of raw byte: \n", response)
		fmt.Println("Printing JSON literal: \n", string(response))

	}

}

func main() {
	servePayloadData()
	fmt.Printf("%x\n", 456)
}
