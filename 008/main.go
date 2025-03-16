// Exercise: User input
// Get a number from the console and check if it's odd
// You can create another function or do it everything in the main func :) 

package main

import "fmt"

func checkInput() {
	var input int
	fmt.Println("Please enter a number:")
	fmt.Scanln(&input)

	if input % 2 == 1 {
		fmt.Printf("%d is odd.", input)
	} else {
		fmt.Printf("%d is even.", input)
	}
}

func main () {
	checkInput()
}