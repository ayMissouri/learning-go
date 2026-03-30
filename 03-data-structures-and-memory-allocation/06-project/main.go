package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	// check if contact already exists
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact %s already exists\n", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact %s added\n", name)
}

func findContact(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("No contacts found")
		return
	}

	for _, contact := range contactList {
		fmt.Printf("%d. %s, %s, %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}

func main() {
	listContacts()

	addContact("John", "john@email.com", "07154885619")
	addContact("Jane", "jane@email.com", "07482165481")
	addContact("Bob", "bob@email.com", "07695484751")
	addContact("John", "john@email.com", "07154885619") // attempt to add duplicate

	listContacts()

	bob := findContact("Bob")
	if bob == nil {
		fmt.Println("Bob not found")
		return
	}
	fmt.Printf("Bob: %v\n", bob)
}
