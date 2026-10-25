package main

import (
	"fmt"
	"go-reloaded/src/convertisseur"
	"go-reloaded/src/decoupage"
	"go-reloaded/src/forme"
	"go-reloaded/src/gestiontxt"
	"go-reloaded/src/ponctuation"
)

func main() {
	a := gestiontxt.Lecture("Sample.txt")
	decoupage.Decoupage(a)
	convertisseur.Convertisseur(decoupage.ListeMot)
	forme.Forme(decoupage.ListeMot)
	decoupage.ListeMot = ponctuation.Ponctuation(decoupage.ListeMot)
	fmt.Println(decoupage.ListeMot)
	b := ponctuation.StringList(decoupage.ListeMot)
	gestiontxt.Ecriture(b)
}
