package atoan

import "go-reloaded/src/decoupage"

func AtoAn() {
	for a := range decoupage.ListeMot {
		if decoupage.ListeMot[a] == "a" {
			if decoupage.ListeMot[a+1][:1] == "a" || decoupage.ListeMot[a+1][:1] == "e" || decoupage.ListeMot[a+1][:1] == "i" || decoupage.ListeMot[a+1][:1] == "o" || decoupage.ListeMot[a+1][:1] == "u" || decoupage.ListeMot[a+1][:1] == "h" {
				decoupage.ListeMot[a] = "an"
			}
		}
	}
}
