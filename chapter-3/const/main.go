package main

const str string = "const string"

const (
	_ = iota
	one
	two
	three
)

func main() {

	println(str)       // string
	str = "new string" // cannot assign to str, will fail in compilation

	println(one)   // 1
	println(two)   // 2
	println(three) // 3
}
