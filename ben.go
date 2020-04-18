package main

import (
	"fmt"
	"math"
)

func main() {

	/* Message principale s'affichant au lancement du programme */

	var choix int

	fmt.Println("WELCOME TO KALKÛL")

	fmt.Println("la calculatrice fontionne comme suit:\n*les operations se font uniquement avec l'entrée de deux nombres\nChoisissez votre numéro pour effectuer l'operation de votre choix")

	fmt.Println("\n1*ADDITION\n2*SOUSTRACTION\n3*MULTIPLICATION\n4*DIVISION\n5*MODULO\n6*RACINE CARRÉE\n7*PARITE D'UN NOMBRE\n")

	fmt.Println("\n\nchoississez votre operation en choisissant un chiffre(1..6)")

	/* variable qui choisis l'operation à effectuer */

	fmt.Scan(&choix)

	/* les différentes operations */
	switch choix {
	/*addition*/
	case 1:
		var a int
		fmt.Printf("Entrez un premier nombre:")
		fmt.Scan(&a)

		var b int
		fmt.Printf("Entrez une second nombre:")
		fmt.Scan(&b)

		resultat := a + b

		fmt.Printf(" %d + %d = %d ", a, b, resultat)
		fmt.Println("")
		break
	/*soustraction*/
	case 2:
		var a int
		fmt.Printf("Entrez un premier nombre:")
		fmt.Scan(&a)

		var b int
		fmt.Printf("Entrez une second nombre:")
		fmt.Scan(&b)

		resultat := a - b

		fmt.Printf(" %d - %d = %d ", a, b, resultat)
		fmt.Println("")
		break
	/*multiplication*/
	case 3:
		var a int
		fmt.Printf("Entrez un premier nombre:")
		fmt.Scan(&a)

		var b int
		fmt.Printf("Entrez une second nombre:")
		fmt.Scan(&b)

		resultat := a * b

		fmt.Printf(" %d x %d = %d ", a, b, resultat)
		fmt.Println("")
		break
	/*division*/
	case 4:
		var a int
		fmt.Printf("Entrez un premier nombre:")
		fmt.Scan(&a)

		var b int
		fmt.Printf("Entrez une second nombre:")
		fmt.Scan(&b)

		resultat := a / b

		fmt.Printf(" %d / %d = %d ", a, b, resultat)
		fmt.Println("")
		break
	/*Modulo*/
	case 5:
		var a int
		fmt.Printf("Entrez un premier nombre:")
		fmt.Scan(&a)

		var b int
		fmt.Printf("Entrez une second nombre:")
		fmt.Scan(&b)

		resultat := a % b

		fmt.Printf(" %d %% %d = %d ", a, b, resultat)
		fmt.Println("")
		break
	/*racine carrée*/
	case 6:
		a := 0.0
		fmt.Printf("Entrez un nombre:")
		fmt.Scan(&a)

		resultat := math.Sqrt(a)

		fmt.Print("√", a, "=", resultat)
		fmt.Println("")
		break
	/*Parité*/
	case 7:
		var a int
		fmt.Printf("Entrer un nombre:")
		fmt.Scan(&a)

		if a%2 == 0 {
			fmt.Printf("le nombre %d est paire", a)
		} else {
			fmt.Printf("le nombre %d est impaire", a)
		}
		fmt.Println("")
		break
	/*message s'affichant en cas d'erreur de choix d'operation*/
	default:
		fmt.Println("error !")
		break
	}

}
