package main

import (
	"encoding/json"
	"fmt"
	//"encoding/xml"
)

type payload struct {
	ParentMap data
}

type data struct {
	Fruit     fruits
	Vegetable vegetables
}

type fruits map[string]int
type vegetables map[string]int

func getJSONPayloadData() ([]byte, error) {
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
	d := data{fruits, vegetables}
	fmt.Printf("%+v\n", d)
	p := payload{d}
	return json.MarshalIndent(p, "", " ")
	//return xml.MarshalIndent(p, "", " ")
}

func servePayloadData() {
	response, err := getJSONPayloadData()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(err)
		fmt.Println(len(response))
		//fmt.Println("Printing slice of raw byte: \n", response)
		//fmt.Println("Printing JSON literal: \n", string(response))
	}

}

func main() {
	servePayloadData()
	fmt.Printf("%x\n", 456)
}
