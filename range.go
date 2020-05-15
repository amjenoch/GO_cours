package main 

import "fmt"

func main() {
	x :=  "Salut"

	for a,b := range x {
		fmt.Println(a,b)
	}
	z :=  [2]int{0,1}

	for a,b := range z {
		fmt.Println(a,b)
	}
	s :=  make(chan int,4)

	s <- 4
	s <- 6
	s <- 9
	s <- 10
	close(s)

	for b := range s {
		fmt.Println(b)
	}
}
