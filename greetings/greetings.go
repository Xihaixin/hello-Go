package greetings
import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == ""{
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hello, %v. Welcome!",
		"Hi, I'm %v. How are you ?",
		"Good morning, %v hahaha.",
		"ennn... %v, what's up?",
	}

	return formats[rand.Intn(len(formats))]
}