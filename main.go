package main

import (
	"fmt"
	"lem-in/parser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error invalid numbre of argument")
		return
	}
	argument := os.Args[1]
	parser.Validation(argument)
}
