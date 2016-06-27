package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// function returns response as a slice of bytes in ASCII decimal code points
// response data code points will need to be converted/casted to string as human readable ASCII characters
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
		fileWriteError := ioutil.WriteFile("sample2.txt", response, 0666)
		if fileWriteError != nil {
			log.Fatal(fileWriteError)
		}
		//fmt.Println(err) prints error value
		fmt.Println("Printing slice of raw byte: \n", response[0:10])
		fmt.Println("Printing slice of raw byte: \n", string(response))
		//fmt.Println("Printing JSON literal: \n", string(response))
		fmt.Println(len(response)) // prints the string length or byte length of the response data
	}
}

func main() {
	servePayloadData()
}
