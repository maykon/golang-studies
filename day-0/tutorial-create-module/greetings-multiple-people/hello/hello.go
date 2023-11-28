package main

import (
	"fmt"

	"log"

	"mk/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    /// A slice of names.
    names := []string{"Maykon", "Érika", "Gabriel"}
    fmt.Println(names)

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}