package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arga := os.Args[1:]
	arg := []byte(arga[0])

	for i := 0; i < len(arg)-1; i++ {
		for j := 0; j < len(arg)-1-i; j++ {
			if arg[j] > arg[j+1] {
				arg[j], arg[j+1] = arg[j+1], arg[j]
			}
		}
	}
	for _, char := range arg {
		z01.PrintRune(rune(char))
		z01.PrintRune('\n')
	}
}
