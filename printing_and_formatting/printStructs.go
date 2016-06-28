package main

import "fmt"

// Project struct
type Project struct {
	ID      int64   `json:"project_id"`
	Title   string  `json:"title"`
	Name    string  `json:"name"`
	Data    Data    `json:"data"`
	Commits Commits `json:"commits"`
}

// Data struct
type Data map[int]int

// Commits struct
type Commits map[int]int

func main() {
	fmt.Printf("%+v\n", Project{}) //To print the name of the fields in a struct
}
