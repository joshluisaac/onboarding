package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Payload struct
type Payload struct {
	ParentMap Data
}

// Data struct
type Data struct {
	Fruit     Fruits
	Vegetable Vegetables
}

// Fruits map
type Fruits map[string]int

// Vegetables map
type Vegetables map[string]int

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
	d := Data{fruits, vegetables}
	p := Payload{d}
	return json.MarshalIndent(p, "", " ")
	//return xml.MarshalIndent(p, "", " ")
}

func servePayloadData(w http.ResponseWriter, r *http.Request) {
	response, err := getJSONPayloadData()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println(err)
		fmt.Println("Printing slice of raw byte: \n", response)
		fmt.Println(len(response))
		fmt.Println("Printing JSON literal: \n", string(response))
		//fmt.Fprintf(w,string(response)) //one way of writing to response writer
		responseInt, _ := w.Write(response) //another way of writing to response writer
		fmt.Println("Total number of bytes: ", responseInt)
	}

}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", servePayloadData)
	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
