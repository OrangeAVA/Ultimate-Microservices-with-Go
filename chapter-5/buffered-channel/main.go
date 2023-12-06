package main

func main() {
	messages := make(chan string, 2)

	messages <- "ping"
	messages <- "pong"

	println(<-messages) // ping
	println(<-messages) // pong
}

func StuckChannel() {
	messages := make(chan string, 1)

	messages <- "ping"
	messages <- "pong"

	println(<-messages)
	println(<-messages)
}
