// Create an array of 5 strings, and initialize it's 2 first values with some random names

package main

import "fmt"

func main () {
	array := [5]string{}
	array[0] = "Madeline"
	array[1] = "Theo"
	
	fmt.Println(array)
}