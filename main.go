package main

import (
	"fmt"
	"go-reloaded/src/gestiontxt"
)

func main() {
	a := gestiontxt.Lecture("Sample.txt")
	fmt.Println(a)
	gestiontxt.Ecriture(a)
}
