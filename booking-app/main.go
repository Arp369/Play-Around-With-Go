package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "still available !!!")
	fmt.Println("Get you tickets here to attend.")
	fmt.Println(conferenceName)
}

// Greet our users
