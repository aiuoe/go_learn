package data

import "fmt"

func Map() {
	var word string

	// maps store values sort and with unique key
	emojis := map[string]string{
		"happy": "😃",
		"smile": "🤣",
		"sad":   "😩",
	}

	emojis["cold"] = "🥶"
	emojis["heat"] = "🥵"

	fmt.Print("find emoji by word: ")
	fmt.Scanln(&word)

	value, ok := emojis[word]

	if ok {
		fmt.Println("emoji:", value)
	} else {
		fmt.Println("this emoji not exists")
	}

	// delete from map by key
	delete(emojis, word)

	for key, value := range emojis {
		fmt.Printf("{key: %s, value: %v}\n", key, value)
	}
}
