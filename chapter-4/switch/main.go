package main

func main() {
	printString("b") // b
	printString("a") // a b
}

func printString(str string) {
	switch str {
	case "a":
		println("a")
		fallthrough
	case "b":
		println("b")
	default:
		println("default")
	}
}
