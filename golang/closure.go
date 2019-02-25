package main

import "fmt"

type Model struct {
	Status int
}

func main() {

	update := func(model *Model) {
		model.Status = 1
	}

	m := &Model{
		Status: 2,
	}
	update(m)

	fmt.Println(m)
}
