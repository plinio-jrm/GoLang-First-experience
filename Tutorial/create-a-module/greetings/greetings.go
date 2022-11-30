package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if (name == "") {
		return "", errors.New("Empty name!") 
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello, %v!!!",
		"Heeeyyyy, you again %v? nice!",
		"Hail, %v! oh my lord!",
	}

	return formats[rand.Intn(len(formats))]
}