package main

type Animal struct {
	Sound string
}

func (a *Animal) Speak() {
	println(a.Sound)
}

type Cat struct {
	Animal // Embedding
}

type Dog struct {
	Animal // Embedding
}

func main() {
	cat := Cat{Animal{"Meow"}}
	cat.Speak() // Meow

	dog := Dog{Animal{"Woof"}}
	dog.Speak() // Woof
}
