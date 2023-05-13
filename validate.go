package main

import "regexp"

func validateUserInput(userTickets int, availableTickets int, err error) (bool, bool, bool, bool) {
	// var isValidName bool = len(firstName) >=2 && len(lastName) >=2
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// Define a regular expression pattern for email validation
	pattern := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	re := regexp.MustCompile(pattern)
	isValidEmail := re.MatchString(email)

	// validate ticket
	isValidTicket := userTickets > 0 && userTickets <= availableTickets && availableTickets > 0
	isValidUserTicket := err == nil

	return isValidName, isValidEmail, isValidTicket, isValidUserTicket
}
