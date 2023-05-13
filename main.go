package main

import (
	"booking-app/greet"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var firstName string
var lastName string
var email string
var wg = sync.WaitGroup{}

const totalTickets int = 50

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets int
}

func main() {

	var again string
	var city string
	var bookings []string
	var mapTypeBookings = make([]map[string]string, 0) // map
	var structTypeBookings = make([]UserData, 0)       // struct

	greet.GreetUsers()

	for greet.Chn_availableTickets > 0 || greet.Bng_availableTickets > 0 && len(bookings) < 50 {

		fmt.Println("------------- CONFERENCE BOOKING ------------")
		fmt.Println("Enter the city where do you want to attend the conference, 'Chennai or Bangalore'")
		fmt.Scan(&city)
		city = strings.ToLower(city)

		switch city {
		case "chennai":

			fmt.Printf("We have total %v tickets out of which %v are available\n", totalTickets, greet.Chn_availableTickets)

			// check ticket sold out status
			ticketSoldOutStatus(greet.Chn_availableTickets)

			// get user input
			chnuserTickets, err := getUserInput()

			// validate user input
			isValidName, isValidEmail, isValidTicket, isValidUserTicket := validateUserInput(chnuserTickets, greet.Chn_availableTickets, err)

			// book tickets
			greet.Chn_availableTickets, bookings, mapTypeBookings, structTypeBookings = ticketBooking(isValidName, isValidEmail, isValidTicket, isValidUserTicket, greet.Chn_availableTickets, chnuserTickets, again, bookings, mapTypeBookings, structTypeBookings)

			wg.Add(1)                                                 // synchronize go routine
			go sendTicket(firstName, lastName, email, chnuserTickets) // go routine for concurrency

		case "bangalore":

			fmt.Printf("We have total %v tickets out of which %v are available\n", totalTickets, greet.Bng_availableTickets)

			// check ticket sold out status
			ticketSoldOutStatus(greet.Bng_availableTickets)

			// get user input
			bnguserTickets, err := getUserInput()

			// validate user input
			isValidName, isValidEmail, isValidTicket, isValidUserTicket := validateUserInput(bnguserTickets, greet.Bng_availableTickets, err)

			// book tickets
			greet.Bng_availableTickets, bookings, mapTypeBookings, structTypeBookings = ticketBooking(isValidName, isValidEmail, isValidTicket, isValidUserTicket, greet.Bng_availableTickets, bnguserTickets, again, bookings, mapTypeBookings, structTypeBookings)

			wg.Add(1)                                                 // synchronize go routine
			go sendTicket(firstName, lastName, email, bnguserTickets) // go routine for concurrency

		default:
			fmt.Println("the city you have entered is invalid..")
		}

		again = continueBooking()
		if again != "y" {
			fmt.Println("See you soon... Have a nice day!")
			wg.Wait()
			os.Exit(0)
		}
		if greet.Chn_availableTickets == 0 && greet.Bng_availableTickets == 0 {
			firstNames := getFirstNames(bookings)
			fmt.Printf("Folks who have purchased tickets in either Chennai or Bangalore. %v\n", firstNames)
			fmt.Printf("map in GO - %v\n", mapTypeBookings)
			fmt.Printf("struct in GO - %v\n", structTypeBookings)
			wg.Wait()
			ticketSoldOutStatus(greet.Chn_availableTickets)
		}
	}

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// println(names[0], names[1])
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (int, error) {

	var userTickets string

	fmt.Println("Enter you firstname:")
	fmt.Scan(&firstName)
	fmt.Println("Enter you lastname:")
	fmt.Scan(&lastName)
	fmt.Println("Enter you email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	num, err := strconv.Atoi(userTickets)
	return num, err
}

func continueBooking() string {
	var again string
	fmt.Println("Do you want to continue booking, press 'y'")
	fmt.Scan(&again)
	return again
}
func ticketSoldOutStatus(availableTickets int) {
	if availableTickets == 0 {
		fmt.Println("Tickets are sold out... Better luck next time!")
		os.Exit(0)
	}
}

func ticketBooking(isValidName bool, isValidEmail bool, isValidTicket bool, isValidUserTicket bool, availableTickets int, userTickets int, again string, bookings []string, mapTypeBookings []map[string]string, structTypeBookings []UserData) (int, []string, []map[string]string, []UserData) {
	if isValidName && isValidEmail && isValidTicket && isValidUserTicket {
		availableTickets = availableTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName) // slice of string

		//create a map - data type
		var allInfo = make(map[string]string)
		allInfo["firstName"] = firstName
		allInfo["lastName"] = lastName
		allInfo["email"] = email
		allInfo["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		mapTypeBookings = append(mapTypeBookings, allInfo) // slice of map

		//create a struct - user defined data type
		var structvar = UserData{
			firstName:   firstName,
			lastName:    lastName,
			email:       email,
			userTickets: userTickets,
		}
		structTypeBookings = append(structTypeBookings, structvar) // slice of struct

		fmt.Printf("Hi, %v %v. Thank you for booking %v tickets. A confirmation email will be sent to %v\n", firstName, lastName, userTickets, email)

	} else {
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email id you have entered is invalid")
		}
		if !isValidUserTicket {
			fmt.Println("Don't try to fool me..! Invalid entry for tickets")
			os.Exit(0)
		}
		if !isValidTicket {
			fmt.Printf("We have only %v tickets available, so you cant book %v tickets\n", availableTickets, userTickets)
		}
	}
	return availableTickets, bookings, mapTypeBookings, structTypeBookings

}

func sendTicket(firstName string, lastName string, email string, userTicket int) {
	time.Sleep(30 * time.Second)
	fmt.Println("################")
	var ticket = fmt.Sprintf("Booking confirmed by %v %v and %v tickets sent to email address %v", firstName, lastName, userTicket, email)
	fmt.Printf("Sending tickets: %v\n", ticket)
	fmt.Println("################")
	wg.Done()
}
