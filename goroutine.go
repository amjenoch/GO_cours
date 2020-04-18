package main 

import (
		"fmt"
		"time"
	)



func f(depuis string)  {
	for i:=0 ; i<4 ; i++{
		fmt.Println(depuis,":",i)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


func main() {


	f("Salut")

	go f("Bonne nuit") 

	go say("world")
	say("hello")

	time.Sleep(time.Second)
    fmt.Println("done")
}


