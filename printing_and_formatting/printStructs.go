package main

import "fmt"

type Project struct {
	Id      int64   `json:"project_id"`
	Title   string  `json:"title"`
	Name    string  `json:"name"`
	Data    Data    `json:"data"`
	Commits Commits `json:"commits"`
}

type Data map[int]int
type Commits map[int]int

func main() {
	fmt.Printf("%+v\n", Project{})
}
