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


func main() {


	f("Salut")

	go f("Bonne nuit") 

	go func(msg string) {
		fmt.Println(msg)
	}("Je veux t'aimer")

	time.Sleep(time.Second)
    fmt.Println("done")
}