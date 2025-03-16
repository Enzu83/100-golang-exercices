// Exercise: Read a file
// Tip: use the "io/ioutil" package

package main

import (
	"fmt"
	"os"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main () {
	f, err := os.Open("read")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())
}