package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error invalid numbre of argument")
		return
	}
	argument := os.Args[1]
	Validation(argument)
}
