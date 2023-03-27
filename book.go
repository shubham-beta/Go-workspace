package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 52
	var remainingTickets uint = 52
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still avaialble.\n", conferenceTickets, remainingTickets, "are still available")
	fmt.Println("Get your tickets here")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastname: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName, lastName)

			fmt.Printf("Thank You %v %v for booking %v tickets. You will receivea confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all of our bookings: %v\n", bookings)

			if remainingTickets == 0 {
				// end program
				fmt.Println(" our conference is booked out. please come back next year")
				break

			} else {
				if !isValidName {
					fmt.Println("first name or last name you entred is too short")
				}

				if !isValidEmail {
					fmt.Println("email address you entered doesn't contain @ sign")

				}
				if !isValidTicketNumber {
					fmt.Println(" number of tickets you've entered is incorrect")
				}

			}

		}

	}

}
