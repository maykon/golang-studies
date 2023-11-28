package main

import (
	"fmt"

	"mk/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Maykon")
    fmt.Println(message)
}