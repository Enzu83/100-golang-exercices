// Exercise: User input
// Using only the fmt package, ask a user for it's name and then for it's surname
// Store it in 2 variables called "name" and "surname"
// After user has entered the data, print it out

// Tip: https://pkg.go.dev/fmt#hdr-Scanning

package main

import "fmt"

func main () {
	var name, surname string

	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Please enter your surname:")
	fmt.Scanln(&surname)

	fmt.Printf("Hello, %s %s!", name, surname)
}
