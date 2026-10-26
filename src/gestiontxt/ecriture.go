package gestiontxt

import (
	"log"
	"os"
)

func Ecriture(s string) {
	err := os.WriteFile("fichiertxt/Result.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
