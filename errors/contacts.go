package _errors

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const FileName string = "contacts.json"

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saver(contacts []Contact) error {
	file, err := os.Create(FileName)

	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func loader(contacts *[]Contact) error {
	file, err := os.Open(FileName)

	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func Contacts() {
	contacts := []Contact{}
	err := loader(&contacts)

	if err != nil {
		log.Println(err)
	}

	var opt int
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(`
		[==> contacts manager <==]
		[1] add contact
		[2] show contacts
		[3] exit 
	
		choose option:
		`)

		_, err := fmt.Scanln(&opt)

		if err != nil {
			log.Println(err)
		}

		switch opt {
		case 1:
			var contact Contact
			fmt.Print("Name: ")
			contact.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			contact.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			contact.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, contact)

			if err := saver(contacts); err != nil {
				log.Println(err)
			}
		case 2:
			fmt.Println("[==============================================================]")
			for index, contact := range contacts {
				fmt.Printf("%d, Name: %s, Email: %s, Phone: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("[==============================================================]")
		case 3:
			return
		default:
			log.Println("invalid option")
		}
	}
}
