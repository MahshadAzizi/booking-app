package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go conference"
	//another way to define a variable and assign it a value
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// uint -> positive integers or name is unsigned integer
	// an array with fixed size
	//var bookings [50]string
	// an array with dynamic size called slice
	var bookings []string

	//fmt.Println("Welcome to", conferenceName, "booking application")
	// print format function
	fmt.Printf("conferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n",
		conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n",
		conferenceTickets, remainingTickets)

	fmt.Println("Get you tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask for their name
		// define variable with pointer, pointer is address of variable in memory
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole slice: %v:\n", bookings)
			fmt.Printf("The first value %v:\n", bookings[0])
			fmt.Printf("slice type %T:\n", bookings[0])
			fmt.Printf("slice length %v:\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confermation email at %v\n",
				firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("the first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("our conference is booked out, com back next year.")
				break
			}
		} else {
			fmt.Printf("we only have %v tickets, you cant book %v tickets\n",
				remainingTickets, userTickets)
		}
	}
}
