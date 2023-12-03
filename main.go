// NOTE: Go programs are organized into packages
// defines the name of the package
package main

//NOTE: Go default packages can be found from the main website at https://pkg.go.dev/std
//NOTE: hovering over the package in your IDE takes you directly to the official documentation of the package
//imports packages from the standard library and also our own custom packages
import (
	"booking-app/helper"

	//This package provides different functions for formatted Input and Output processing functionalities
	"fmt"
	"sync"
	"time"
)

// global variables declaration
// can be used anywhere by any function within the program
const conferenceTickets int = 50

// NOTE: The var keyword is used to declare variables
// NOTE: Go has Type inference and therefore infers data types when declaring variables
// NOTE: Type inference applies to variables declared within functions but for global variables the variable data type should be defined
// declares a variable conferenceName and assigns it the value 'GO Conference"
// NOTE: This is an example of a package level variable which means it is declared outside all functions it can be used anywhere within the same package
// add value data type using the string keyword
var conferenceName string = "Go Conference"

var remainingTickets uint = 50

// this is a map
// it is used to store key value pairs
// the zero indicates the initial value of the map which ideally can be 1 or 0
var bookings = make([]userData, 0)

// NOTE: The type keyword is used to create a new type with the name you have given to your structure
// struct used to declare the basic structure of the expected dataset and its types
type userData struct {
	firstName string
	lastName  string
	email     string

	//uint is an integer type whose value cannot be negative
	numberOfTickets uint
}

// A waitgroup is used to check and wait for launched goroutine to finish
// Sync provides basic synchronization functionality
var wg = sync.WaitGroup{}

// NOTE: To run a go program in the terminal use the command: go run [name of the file] e.g go run main.go
// NOTE: To turn this into an executable file run the command: go build [name of the file] e.g go build main.go
// This is the programs starting/entry point
// an application can have only one entry point
func main() {
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	if isValidName && isValidEmail && isValidTicketNumber {

		//booking method
		bookTicket(userTickets, firstName, lastName, email)

		//specifies number of go routines that the main thread should wait for before creating new threads
		wg.Add(1)

		// Instruct the program to break this process from the main thread into its own thread and continue with the rest of the program below
		// Then depending on the delay time set on the function declaration the execution will execute after the set time
		//The 'go' keyword starts of a new goroutine ( def: goroutine is a lightweight thread managed by the GO runtime)
		go sendTickets(userTickets, firstName, lastName, email)

		//NOTE: This applies to variables crated within functions aka scoped variables for global variables the data type should be stated during declaration
		//NOTE: This type of declaration does not apply to variable declarations where the variable value is a const value
		//variable assignment shorthand with infered data types
		firstNames := getFirstNames()

		//Printf enables us to print formatted strings aka strings with variable values within them
		//%v aka place holder is used to show the position in the string where the variable value should be passed
		fmt.Printf("The bookings available currently are: %v\n", firstNames)

		if remainingTickets == 0 {
			//prints a string and pushes the cursor to the next line
			fmt.Println("Our conference is booked out. Come back next year ")
		}
	} else {
		if !isValidName {
			fmt.Println("Firstname or lastname you entered is too short")
		}

		if !isValidEmail {
			fmt.Println("Email address you entered is not valid")
		}

		if isValidTicketNumber {
			fmt.Println("The number of tickets you entered is invalid")
		}
	}

	//Another waitgroup function usually placed at the end of the function to wait for all threads that were queued to execute before exiting the program
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and only %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Purchase your tickets now to attend ")
}

func getFirstNames() []string {
	//NOTE: A slice is different from an array in that it is dynamic in size i.e its size increases as values are added
	//defines a slice to store the first names of all bookings
	firstNames := []string{}

	//NOTE: Range is used to iterate over elements and in this case returns the index and value of each element
	//NOTE: For values that you don't intend to use in the for loop they are shown using the '_' aka Blank Identifier (This helps solve the declared but not used error for such)
	//foreach loop to loop through all bookings and append the firstname value to the firstnames slice
	for _, booking := range bookings {
		//append function is used to add elements to a slice
		//It takes the slice and value as parameters
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	//NOTE: To store a variable value in memory the pointer is used
	//NOTE: A pointer is denoted using the '&' sign
	//NOTE: In Go a pointer is also known as a special variable
	//used to get and store firstname value entered by the user into the firstName variable
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	//NOTE: A local variable can only be used within that function
	//This is an example of a local variable declared within a function
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\n ", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets Remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("List of bookings is %v\n", bookings)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	//This is a time that basically delays execution of the logic in this function/("thread") for 30 seconds
	time.Sleep(30 * time.Second)

	//The Sprintf function from fmt package helps to create and as well save the created string into a variable
	//NOTE: Trying to do this using the regular Printf is not possible
	var ticket = fmt.Sprintf("%v tickets for %v, %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v\n to email address %v\n", ticket, email)
	fmt.Println("################")

	//Another waitgroup function that removes a thread from the waitgroup waiting list
	//Signifies that executing of all threads is complete
	wg.Done()
}
