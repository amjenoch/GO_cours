/*package main

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
*/

package main

import "fmt"

func ordre(p []string) []string {

	for i := 0; i < len(p); i++ {
		for k, l := range p {
			if p[i] < l {
				p[k], p[i] = p[i], p[k]
			}
		}
	}
	return p
}
func main() {

	m, s, j, z, b, y, c := "", "", "", "", "", "", ""

	fmt.Println("Vous devez entrer 7 lettres...\n")

	fmt.Print("Entrer la première lettre: ")
	fmt.Scan(&m)
	fmt.Print("Entrer la seconde lettre: ")
	fmt.Scan(&s)
	fmt.Print("Entrer la troisième lettre: ")
	fmt.Scan(&j)
	fmt.Print("Entrer la quatrième lettre: ")
	fmt.Scan(&z)
	fmt.Print("Entrer la cinquième lettre: ")
	fmt.Scan(&b)
	fmt.Print("Entrer la sixième lettre: ")
	fmt.Scan(&y)
	fmt.Print("Entrer la septième lettre: ")
	fmt.Scan(&c)

	p := []string{m, s, j, z, b, y, c}
	fmt.Println(ordre(p))
}
