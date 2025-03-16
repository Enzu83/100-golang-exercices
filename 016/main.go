// Exercise: Check if a file exists
// Tip: use the "os" package

package main

import (
	"fmt"
	"os"
)


func main () {
	_, err := os.ReadFile("file")
	
	if err == nil {
		fmt.Println("File found.")
	} else {
		fmt.Println("File not found.")
	}
}