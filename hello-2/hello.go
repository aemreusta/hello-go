package main

import (
	"example/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Emre")
	fmt.Println(message)
}
