package main

import (
	"fmt"
)

func echo() (string, int, int) {
	fmt.Println("")
	return "d", 4, 0
}

func main() {
	a,b,c := echo()
	fmt.Println(a,b,c)
}