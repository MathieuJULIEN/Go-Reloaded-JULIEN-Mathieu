package ponctuation

import (
	"go-reloaded/src/decoupage"
	"strings"
)

func CestPonctuation(mot string) bool {
	if len(mot) == 1 {
		if mot == "," || mot == "." || mot == "!" || mot == "?" || mot == ":" || mot == ";" {
			return true
		}
	} else {
		for a := range mot {
			if string(mot[a]) != "," && string(mot[a]) != "." && string(mot[a]) != "!" && string(mot[a]) != "?" && string(mot[a]) != ":" && string(mot[a]) != ";" {
				if a == len(mot)-1 {
					return false
				}
			} else {
				return true
			}
		}
	}
	return false
}

func Ponctuation() {
	for a := 0; a < len(decoupage.ListeMot); a++ {
		if CestPonctuation(decoupage.ListeMot[a]) && a > 0 {
			decoupage.ListeMot[a-1] += decoupage.ListeMot[a]
			decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
			a--
		}
	}
}

func StringList(list []string) string {
	return strings.Join(list, " ")
}
