/*
planning to conduct a standup comedy show in my home town
so i dicided to create an appliation with a simple command prompt application
no ticket cancelation as there is less time we welcome you to book your tickets
as soon as possible
*/
package main

import (
	"CLI_Application/Ticketsent" //ticket to send to customer module
	"CLI_Application/booking"    //booking ticket module
	"CLI_Application/client"     //getting client data
	"CLI_Application/greeting"   //welcoming the uset
	"CLI_Application/helper"     //validating the user details
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	conferenceName   = "Stand-up comedy"
	remainingTickets uint
	comedian         = "Zakir-khan"
	sum              uint
)

const conferenceTickets uint = 250

func main() {

	db, err := sql.Open("mysql", "root:a@11189D001#@tcp(127.0.0.1:3306)/mydata")
	if err != nil {
		fmt.Println("error in connecting the db", err)
	}
	err = db.QueryRow("SELECT SUM(userTickets) FROM customer").Scan(&sum)
	if err != nil {
		fmt.Println("error in query row statement", err)
		return
	}

	defer db.Close()

	remainingTickets = conferenceTickets - sum

	greeting.GreetUsers(comedian, conferenceName, remainingTickets, sum)

	firstName, lastName, email, userTickets := client.GetUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	stmt, err := db.Prepare("INSERT INTO customer(firstName, lastName, email, userTickets) values(?,?,?,?)")
	if err != nil {
		fmt.Println("error in statement object", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(firstName, lastName, email, userTickets)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isValidName && isValidEmail && isValidTicketNumber {

		booking.BookTicket(userTickets, firstName, lastName, email, conferenceName)

		Ticketsent.SendTicket(userTickets, firstName, lastName, email)

		firstNames := booking.GetFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}

		fmt.Println(`have a nice day! keep smiling keep shining $`)

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
}
