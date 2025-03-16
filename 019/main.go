// Exercise: Rename a file from name1 to name2
// Tip: use the "os" package

package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var src, dest string
	src = "name1.txt"
	dest = "name2.txt"

	err := os.Rename(src, dest)
	check(err)
}