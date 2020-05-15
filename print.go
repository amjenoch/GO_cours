package main

import "fmt"

func uniqueNames(a, b []string) []string {
	var result []string
	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
