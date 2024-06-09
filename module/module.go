package gomodule

import (
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is required")
	}

	return fmt.Sprintf(randonMessage(), name), nil
}

func Greeter(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randonMessage() string {
	message := []string{
		"Hello, %s Welcome!\n",
		"goodbye %s",
		"%s, you are welcome",
	}

	return message[rand.Intn(len(message))]
}
