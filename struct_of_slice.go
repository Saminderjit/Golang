package main

import (
	"fmt"
)

type person struct {
	name     string
	lastName string
	age      int
}

type people struct {
	numberOfpeople []person
}

func main() {
	fmt.Println("Hello, playground")
	saminder := people{}

	saminder = people{numberOfpeople: []person{person{name: "Saminder", lastName: "Chahal", age: 25}}}

	fmt.Println("%v\n", saminder)

}
