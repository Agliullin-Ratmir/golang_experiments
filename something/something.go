package something

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

type UserData struct {
	firstName string
}

func DoSomething() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome ", conferenceName)
	fmt.Println("Total tickets ", conferenceTickets)
	fmt.Println("Remaining tickets ", remainingTickets)

	var firstName string
	var secondName string
	var email string
	var userTickets uint

	fmt.Printf("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter second name: ")
	fmt.Scan(&secondName)

	fmt.Printf("Enter email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter amount of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	var bookings []string
	bookings = append(bookings, firstName+" "+secondName)

	fmt.Printf("the whole array: %v", bookings)
	fmt.Printf("the first element: %v", bookings[0])

	fmt.Printf("Thank you, %v for %v tickets", firstName, userTickets)
	helper.PrintRemainTickets(remainingTickets)

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		fmt.Println(names[0])

		var noTicketsRemaining bool = remainingTickets == 0
		if noTicketsRemaining {

		} else {

		}
	}
}
