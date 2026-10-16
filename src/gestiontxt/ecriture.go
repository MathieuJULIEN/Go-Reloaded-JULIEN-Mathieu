package gestiontxt

import (
	"log"
	"os"
)

func Ecriture(s string) {
	err := os.WriteFile("src/fichiertxt/Result.txt", []byte(s), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
