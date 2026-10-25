package ponctuation

import "strings"

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

func Ponctuation(liste []string) []string {
	for a := 0; a < len(liste); a++ {
		if CestPonctuation(liste[a]) && a > 0 {
			liste[a-1] += liste[a]
			liste = append(liste[:a], liste[a+1:]...)
			a--
		}
	}
	return liste
}

func StringList(list []string) string {
	return strings.Join(list, " ")
}
