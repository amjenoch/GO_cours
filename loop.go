package main

import "fmt"

/*func main() {
	for i:=0 ; i<10; i++{
		for j:=0 ; j<3 ; j++{
			fmt.Printf("For i=%d\t on a j=%d\n",i,j)
		}
	}
}
*/

func main() {
	for i:=0 ; i<10; i++{
		fmt.Printf("For i=%d\n",i)
		for j:=0 ; j<3 ; j++{
			fmt.Printf("\t\t on a j=%d\n",j)
		}
	}
}