package main

import (
	"fmt"
	"go-reloaded/src/atoan"
	"go-reloaded/src/convertisseur"
	"go-reloaded/src/decoupage"
	"go-reloaded/src/forme"
	"go-reloaded/src/gestiontxt"
	"go-reloaded/src/ponctuation"
)

func main() {
	a := gestiontxt.Lecture("Sample.txt")
	decoupage.Decoupage(a)
	convertisseur.Convertisseur()
	forme.Forme()
	ponctuation.Ponctuation()
	fmt.Println(decoupage.ListeMot)
	atoan.AtoAn()
	fmt.Println(decoupage.ListeMot)
	b := ponctuation.StringList(decoupage.ListeMot)
	gestiontxt.Ecriture(b)
}
