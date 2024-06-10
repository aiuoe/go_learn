package gomodule

import (
	"fmt"
	"log"
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

func Exec() {
	greeting, err := Hello("rub3n")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(greeting)

	names := []string{"rub3n", "adriana", "keiler"}

	messages, err := Greeter(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		log.Println(message)
	}
}
