package helper

import "strings"

// NOTE: This is declared outside all functions and the first letter is to be capitalized
// NOTE: This variable can be used everywhere across all packages
// This os an example of a Global Variable
var AnotherVar = "value"

// NOTE: To enable a function to be 'exportable' or rather be able to be imported into another package the first letter of the
// function name is capitalized. In this can the 'V' in ValidateUserInput is made Capital
// NOTE: Same concept of capitalizing functions to make them exportable applies to variables that exist into the go package file you intend to be exportable
// This is a function that validates user inputs and can be imported into any other go file as a package
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
