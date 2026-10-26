package ponctuation

import (
	"go-reloaded/src/decoupage"
	"strings"
)

func CestPonctuation(mot string) (bool, int) {
	var nb int = 0
	if len(mot) == 1 {
		if mot == "," || mot == "." || mot == "!" || mot == "?" || mot == ":" || mot == ";" {
			return true, 1
		}
	} else {
		for a := range mot {
			if string(mot[a]) != "," && string(mot[a]) != "." && string(mot[a]) != "!" && string(mot[a]) != "?" && string(mot[a]) != ":" && string(mot[a]) != ";" {
				if a == 0 {
					return false, 0
				} else {
					return true, nb
				}
			} else {
				nb++
				if a == len(mot)-1 {
					return true, nb
				}
			}
		}
	}
	return false, 0
}

func Ponctuation() {
	for a := 0; a < len(decoupage.ListeMot); a++ {
		tempo, nb := CestPonctuation(decoupage.ListeMot[a])
		if tempo == true && a > 0 {
			decoupage.ListeMot[a-1] += decoupage.ListeMot[a][:nb]
			decoupage.ListeMot[a] = decoupage.ListeMot[a][nb:]
			decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a:]...)
			a--
		}
	}
}

func StringList(list []string) string {
	return strings.Join(list, " ")
}
