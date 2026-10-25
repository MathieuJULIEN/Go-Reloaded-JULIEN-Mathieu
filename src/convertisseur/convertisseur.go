package convertisseur

import (
	"fmt"
	"go-reloaded/src/decoupage"
	"strconv"
)

func Convertisseur(liste []string) {
	for a := range liste {
		if liste[a] == "(hex)" {
			provi, err := strconv.ParseInt(liste[a-1], 16, 64)
			if err == nil {
				decoupage.ListeMot[a-1] = strconv.FormatInt(provi, 10)
				decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
				a--
			} else {
				fmt.Println("Erreur: Ce qu'il y a avant le (hex) n'est pas de l'hexadecimal")
			}
		} else if liste[a] == "(bin)" {
			provi, err := strconv.ParseInt(liste[a-1], 2, 64)
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
