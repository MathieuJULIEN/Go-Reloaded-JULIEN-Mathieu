package main

import (
	"go-reloaded/src/apostrophe"
	"go-reloaded/src/atoan"
	"go-reloaded/src/convertisseur"
	"go-reloaded/src/decoupage"
	"go-reloaded/src/forme"
	"go-reloaded/src/gestiontxt"
	"go-reloaded/src/ponctuation"
	"go-reloaded/src/saisifichier"
)

func main() {
	saisi := saisifichier.SaisiFichier()
	a := gestiontxt.Lecture(saisi)
	decoupage.Decoupage(a)
	convertisseur.Convertisseur()
	forme.Forme()
	ponctuation.Ponctuation()
	atoan.AtoAn()
	apostrophe.Apostrophe()
	b := ponctuation.StringList(decoupage.ListeMot)
	gestiontxt.Ecriture(b)
}
