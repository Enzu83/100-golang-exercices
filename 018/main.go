// Exercise: Write a list of 5 countries to a file
// Tip: use the "os" package

package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("write")
	check(err)

	countries := []string{"France", "Germany", "Spain", "United Kingdom", "Japan"}

	for _, country := range countries {
		f.WriteString(country + "\n")
	}
}