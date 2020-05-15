package main

import "fmt"

	const (
		_ = iota 
		x = 1 << (iota * 3)
		y = 1 << (iota * 2)
	)


func main() {
	fmt.Printf("%b\n",x)
	fmt.Printf("%b",y)
}