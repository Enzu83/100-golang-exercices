// Exercise: Concatenate two string variables and print it's result
// One variable should be called "helloWorld" and the other "itsmemario"
// There should be a space between them

package main

import "fmt"

func main() {
	var helloWorld, itsmemario string
	
	helloWorld = "Hello world! "
	itsmemario = "It's-a-me, Mario! Waouh!"

	fmt.Println(helloWorld + itsmemario)
}
