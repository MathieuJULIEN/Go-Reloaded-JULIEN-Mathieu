package forme

import (
	"go-reloaded/src/decoupage"
	"strings"
)

func Forme(liste []string) {
	for a := 0; a < len(decoupage.ListeMot); a++ {
		if decoupage.ListeMot[a] == "(up)" {
			provi := strings.ToUpper(liste[a-1])
			decoupage.ListeMot[a-1] = provi
			decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
		} else if decoupage.ListeMot[a] == "(low)" {
			provi := strings.ToLower(liste[a-1])
			decoupage.ListeMot[a-1] = provi
			decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
		} else if decoupage.ListeMot[a] == "(cap)" {
			provi := strings.Title(liste[a-1])
			decoupage.ListeMot[a-1] = provi
			decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
		}
	}
}
