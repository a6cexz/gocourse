package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var addressBook map[string][]int

func loadAddressBook() {
	if addressBook == nil {
		data, err := ioutil.ReadFile("book.json")
		if os.IsNotExist(err) {
			addressBook = make(map[string][]int)
			return
		}

		if err != nil {
			fmt.Printf("Error reading address book %v\n", err)
			return
		}

		err = json.Unmarshal(data, &addressBook)
		if err != nil {
			fmt.Printf("Error parsing address book json %v\n", err)
			return
		}
	}
}

func saveAddressBook() {
	if addressBook == nil {
		return
	}

	data, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Printf("Error creating address book json %v\n", err)
		return
	}

	err = ioutil.WriteFile("book.json", data, 0644)
	if err != nil {
		fmt.Printf("Error creating address book json %v\n", err)
		return
	}
}

// Add adds new person to address book
func Add(name string, phone int) {
	loadAddressBook()
	var phones []int
	if phones, ok := addressBook[name]; !ok {
		phones = make([]int, 1)
		addressBook[name] = phones
	}
	phones = append(phones, phone)
	addressBook[name] = phones
	saveAddressBook()
}

// Print prints address book
func Print() {
	loadAddressBook()
	if addressBook == nil || len(addressBook) == 0 {
		fmt.Println("Address book is empty")
		return
	}

	fmt.Println("Address Book:")
	for name, phones := range addressBook {
		fmt.Printf("Name: %v\n", name)
		for i, phone := range phones {
			fmt.Printf("\t %v) %v\n", i+1, phone)
		}
	}
}

// Run runs address book demo
func Run() {
	fmt.Println()
	fmt.Println("Address Book Demo")
	fmt.Println()

	Print()
	fmt.Println()

	fmt.Println("Adding items")
	Add("Alex", 123789)
	Add("Stan", 987232)
	Add("Steve", 456789)
	Add("Alan", 1789)
	Add("Alex", 456)

	Print()
}
