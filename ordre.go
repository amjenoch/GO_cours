package main

import "fmt"

func ordre(p []string) []string {

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
	fmt.Println(ordre([]string{"m", "s", "j", "z", "b", "y", "c"}))
}