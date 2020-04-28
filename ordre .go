/*package main

import "fmt"

func pnange(p []string) []string  {

	 tab := []string{}
	for i:=0 ; i <len(p);i++{
		p[i],p[0] = p[0],p[i]
		tab = append(tab,p[i])
	}
	return tab
}

func main() {

	fmt.Println(pnange([]string{"l","h","e","c"}))
}*/

/*package main

import "fmt"

func pnange(p []int) []int  {

	 tab := []int{}
	for i:=0 ; i <len(p);i++{
		for k,l := range p {
			if p[i] < l {
				p[k],p[i] = p[i],p[k]
				tab = append(p)
			}
		}
	}
	return tab
}

func main() {

	fmt.Println(pnange([]int{5,7,9,4}))
}
*/

/*package main

import (
	"fmt"
)

func Abort(a, b, c, d, e string) string {
	Tab := []string{a, b, c, d, e}

	for i:=0; i<len(Tab);i++{
		for k,l := range Tab {
			if Tab[i] < l {
				Tab[k],Tab[i] = Tab[i],Tab[k]

			}
		}
	}
	return Tab[len(Tab)/2]

}



func main() {
	middle := Abort("c", "e", "d", "a", "b")
	fmt.Println(middle)
}*/

package main

import "fmt"

func pnange(p []string) []string {

	tab := []string{}
	for i := 0; i < len(p); i++ {
		for k, l := range p {
			if p[i] < l {
				p[k], p[i] = p[i], p[k]
				tab = append(p)
			}
		}
	}
	return tab
}

func main() {

	fmt.Println(pnange([]string{"z", "b", "y", "c"}))
}
