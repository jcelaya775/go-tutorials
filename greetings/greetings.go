package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// NOTE: Go functions can return multiple values
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

func randomFormat() string {
	// NOTE: A slice is a resizable list
	// A slice of messag formats
	// array of strings: [n]string, where n is the length of the array
	// slice of strings: []string
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
