package main

import "fmt"


func affichage(p *int) {
	fmt.Println("Valeur de:",*p)
}

func calc(p *int) {
	*p = *p + 4
}


func main() {

	var a int = 4
	var b *int = &a
 
	affichage(&a)
	calc(b)
	affichage(&a)
}