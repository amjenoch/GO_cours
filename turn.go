package main 

import (
	"fmt"
	"math"
	"os"
)


	var array = []float32{0.1,0.2,0.3} //const



	const (
		a = iota 
		b = iota 
		c = iota
	)

	const (
		d,e,f = iota,iota,iota 
	)

		const (
			Pi = math.Pi
		)



func main() {

	s := append([]byte("hello"),"world"...)
		x:=1
	fmt.Printf("%v\n",array)
	fmt.Println(x)
	fmt.Println(a,b,c,d,e,f)
	fmt.Println(s)
	fmt.Println(Pi)
	tagsViews := map[string]int{
		"unix": 0,
		"python": 1,
		"go": 2,
		"golang": 3,
		"devops": 4,
		"gc": 5,
	}
	for key,views := range tagsViews{
		fmt.Println("There are",views,"views for",key)
	}

	src, err := os.Open("filename")
	// defer src.Close()

	if err != nil {
		return
	}

	// defer src.Close()

	src.WriteString("hello")

	// defer src.Close()

	

}