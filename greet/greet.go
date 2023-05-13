package greet

import (
	"fmt"
)

var conferenceName string = "Nullcon Conference"
var Bng_availableTickets int = 50
var Chn_availableTickets int = 50

func GreetUsers() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("In both Chennai and Bangalore, there is a combined total of %v tickets available\n", Bng_availableTickets+Chn_availableTickets)
	fmt.Printf("Get your tickets here...\n\n")
}
