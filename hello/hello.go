package main

import (
	"fmt"
	"log"
	"strings"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Garrin"}

	for i := 0; i < 10; i++ {
		// Request a greetings message.
		messages, err := greetings.Hellos(names)
		if err != nil {
			log.Fatal(err)
		}
		// If no error was returned, print the returned map of
		// messages to the console.
		fmt.Println(messages)
	}

  x := [4]int{1, 2, 3, 4}
  fmt.Printf("%T\n", x)
  fmt.Println(x)
}
