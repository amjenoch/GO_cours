package main 

import (
	"fmt"
)
/*
func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	fmt.Println(<- messages)
}
*/


/*package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y,x-y)
}
*/



/*package main

import "fmt"

func sum(d []int, e chan int) {
	step := 0 

	for i:= 0; i<len(d);i++{
		 step += i
	}
	e <- step
}



func main() {
	s :=  []int{7,8,9,44,2}

	c := make(chan int)

	go sum(s[len(s)/2:],c)
	go sum(s[:len(s)/2],c)

	x,y := <- c, <- c 

	fmt.Println(x,y,x-y)

	
}*/


func main() {
	c:= make(chan int)

	go func() {
		c <- 42
	}()

	x := <- c 

	fmt.Println(x)
}