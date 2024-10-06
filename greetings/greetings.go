package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.

func Hello(name string) (string,error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("Empty Name!!!")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.


	message := fmt.Sprintf("Hi, %v. Welcome!, And this is your first Go example", name)
	return message, nil
}
