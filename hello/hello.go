package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // seta a flag to disable printing the time, source file, and line number

	for i := 0; i < 10; i++ {
		message, err := greetings.Hello("Gladys")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message)
	}
}
