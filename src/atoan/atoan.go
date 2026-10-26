package atoan

import "go-reloaded/src/decoupage"

func AtoAn() {
	for a := range decoupage.ListeMot {
		if decoupage.ListeMot[a] == "a" {
			if decoupage.ListeMot[a+1][:0] == "a" || decoupage.ListeMot[a+1][:0] == "e" || decoupage.ListeMot[a+1][:0] == "i" || decoupage.ListeMot[a+1][:0] == "o" || decoupage.ListeMot[a+1][:0] == "u" || decoupage.ListeMot[a+1][:0] == "h" {
				decoupage.ListeMot[a] = "an"
			}
		}
	}
}
