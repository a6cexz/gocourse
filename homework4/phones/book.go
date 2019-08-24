package phones

import (
	"fmt"
	"sort"
)

// Contact interface
type Contact interface {
	GetName() string
	SetName(value string)

	GetPhone() int
	SetPhone(value int)
}

// PhoneBook interface
type PhoneBook interface {
	Sort()
	AddContact(c *Contact)
	Contacts() []Contact
}

// ContactData struct
type ContactData struct {
	name  string
	phone int
}

// GetName returns contact name
func (c *ContactData) GetName() string {
	return c.name
}

// SetName sets contact name
func (c *ContactData) SetName(value string) {
	c.name = value
}

// GetPhone returns contact phone
func (c *ContactData) GetPhone() int {
	return c.phone
}

// SetPhone returns contact phone
func (c *ContactData) SetPhone(value int) {
	c.phone = value
}

// PhoneBookData stores contacts
type PhoneBookData struct {
	contacts []Contact
}

// Sort function sorts contacts
func (b *PhoneBookData) Sort() {
	less := func(i, j int) bool {
		return b.contacts[i].GetName() < b.contacts[j].GetName()
	}
	sort.Slice(b.contacts, less)
}

// AddContact adds new contact
func (b *PhoneBookData) AddContact(c Contact) {
	b.contacts = append(b.contacts, c)
}

// Contacts returns stored contacts
func (b *PhoneBookData) Contacts() []Contact {
	return b.contacts
}

// Run functions shows sort demo
func Run() {
	fmt.Println()
	fmt.Println("Phones Demo")
	fmt.Println()

	book := &PhoneBookData{}

	c1 := &ContactData{}
	c1.SetName("John")
	c1.SetPhone(123987)
	book.AddContact(c1)

	c2 := &ContactData{}
	c2.SetName("Victor")
	c2.SetPhone(245678)
	book.AddContact(c2)

	c3 := &ContactData{}
	c3.SetName("Alexey")
	c3.SetPhone(3465182)
	book.AddContact(c3)

	book.Sort()

	for _, c := range book.Contacts() {
		fmt.Printf("%v: %v\n", c.GetName(), c.GetPhone())
	}
}
