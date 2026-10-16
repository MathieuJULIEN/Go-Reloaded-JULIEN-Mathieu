package main

import (
	"fmt"
	"go-reloaded/src/decoupage"
	"go-reloaded/src/lecture"
)

func main() {
	a := lecture.Lecture("Sample.txt")
	fmt.Println(a)
	decoupage.Decoupage(a)
}
