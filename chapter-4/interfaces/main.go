package main

import "fmt"

type Person struct {
	Name string
}

func (m Person) Print() {
	fmt.Println(m.Name)
}

type DogPerson struct {
	Name string
}

type PrintInterface interface {
	Print()
}

var PrintableObject PrintInterface = Person{}

// var PrintableObject2 PrintInterface = DogPerson{} // This will throw an error in compile time

func main() {

}
