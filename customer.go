package main

import "fmt"

//Declare a struct
type customer struct {
	firstName string
	lastName string
	username string
	password string
	email string
	phone int
	address string
}

//Create Methods
func (customer customer) retrieveCredentials() (string, string) {
	return customer.username, customer.password
}

func (customer customer) retrieveAddress() string {
	return customer.address
}

func (customer customer) printAll() {
	fmt.Printf("First Name: %v\n", customer.firstName)
	fmt.Printf("Last Name: %v\n", customer.lastName)
	fmt.Printf("Username: %v\n", customer.username)
	fmt.Printf("Password: %v\n", customer.password)
	fmt.Printf("Email: %v\n", customer.email)
	fmt.Printf("Phone: %v\n", customer.phone)
	fmt.Printf("Address: %v\n", customer.address)
	fmt.Println("")
}