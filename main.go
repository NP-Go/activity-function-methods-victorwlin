package main

import "fmt"

func main() {
	//Declare new item customer
	mj := customer{"Michael", "Jordan", "MJ2020", "1234567", "MJ2020@gmail.com", 12345678, "18227 Capstan Greens Road Cornelius, NC 28031"}

	//Insert Code here
	mj.printAll()

	username, password := mj.retrieveCredentials()
	fmt.Printf("Username: %v\n", username)
	fmt.Printf("Password: %v\n", password)
	fmt.Println("")

	address := mj.retrieveAddress()
	fmt.Printf("Address: %v\n", address)
	fmt.Println("")
}
