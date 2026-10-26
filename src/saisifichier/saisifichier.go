package saisifichier

import "fmt"

var NomFichier string

func SaisiFichier() string {
	fmt.Println("Saisir le nom du fichier.txt (ne pas oublier le .txt et les majuscules au bon endroit): ")
	fmt.Scanln(&NomFichier)
	return NomFichier
}
