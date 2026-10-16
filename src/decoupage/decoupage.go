package decoupage

import (
	"log"
	"os"
)

func Decoupage(s string) {
	err := os.WriteFile("src/fichiertxt/Result.txt", []byte(s), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
