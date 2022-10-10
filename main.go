/*
QUESTION:
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered. 
Truncation is the process of removing the digits to the right of the decimal place.
This code is my solution for a coursera problem.
*/
package main

import ( //imports the format and log
	"fmt"
	"log"
)

var inputN float64

func main() {
	fmt.Println("Please enter a Float number:") // gets user input 
	_, err := fmt.Scan(&inputN) // uses scan to read the users input 
	if err != nil { 
		log.Printf("[Error] Invalid input !")
	}

	fmt.Printf("The number you've entered is '%v'.\n", inputN)
	fmt.Printf("Truncated version of '%v' is '%v'.\n", inputN, int64(inputN))
}