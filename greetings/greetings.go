package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// NOTE: Go functions can return multiple values
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Another way to declare a variable:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retreived message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// NOTE: A slice is a resizable list
	// A slice of message formats
	// array of strings: [n]string, where n is the length of the array
	// slice of strings: []string
	// A slice are like references to an array
	formats := []string{ // This creates an array literal of [3]string, and then builds a slice that references it
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
