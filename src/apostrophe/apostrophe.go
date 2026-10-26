package apostrophe

import "go-reloaded/src/decoupage"

func Apostrophe() {
	var compte int = 0
	for a := 0; a < len(decoupage.ListeMot); a++ {
		if decoupage.ListeMot[a] == "'" {
			compte += 1
			if compte == 1 {
				decoupage.ListeMot[a+1] = "'" + decoupage.ListeMot[a+1]
				decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
			} else if compte == 2 {
				compte = 0
				decoupage.ListeMot[a-1] = decoupage.ListeMot[a-1] + "'"
				decoupage.ListeMot = append(decoupage.ListeMot[:a], decoupage.ListeMot[a+1:]...)
			}
		}
	}
}
