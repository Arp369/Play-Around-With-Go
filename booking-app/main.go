package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available !!! \n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend.")

	var userName string
	var userTickets int
	var userEmail string

	fmt.Print("Enter your name:")
	fmt.Scan(&userName)
	fmt.Print("Enter your valid email:")
	fmt.Scan(&userEmail)
	fmt.Print("Enter the numbers of tickets you want:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation message on %v shortly!!", userName, userTickets, userEmail)
}

// Greet our users
