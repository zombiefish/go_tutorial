package main

import {
	"fmt"
	"github.com/zombiefish/go_tutorial/blob/master/greetings/greetings.go"
}

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.PrintIn(message)
}
