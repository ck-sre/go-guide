package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "how", "are", "you"}
	updateSlice(mySlice) // Original slice gets changed.
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
