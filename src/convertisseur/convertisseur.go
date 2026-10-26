package convertisseur

import (
	"fmt"
	"go-reloaded/src/decoupage"
	"strconv"
)

func Convertisseur() {
	for a := 0; a < len(decoupage.ListeMot); a++ {
		if decoupage.ListeMot[a] == "(hex)" {
			provi, err := strconv.ParseInt(decoupage.ListeMot[a-1], 16, 64)
			if err == nil {
				decoupage.ListeMot[a-1] = strconv.FormatInt(provi, 10)
				decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
				a--
			} else {
				fmt.Println("Erreur: Ce qu'il y a avant le (hex) n'est pas de l'hexadecimal")
			}
		} else if decoupage.ListeMot[a] == "(bin)" {
			provi, err := strconv.ParseInt(decoupage.ListeMot[a-1], 2, 64)
			if err == nil {
				decoupage.ListeMot[a-1] = strconv.FormatInt(provi, 10)
				decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
				a--
			} else {
				fmt.Println("Erreur: Ce qu'il y a avant le (bin) n'est pas du binaire")
			}
		}
	}
}
