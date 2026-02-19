package main


import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error invalid numbre of argument")
		return
	}
	argument  := os.Args[1]
	file , err := os.ReadFile(argument)
	if err != nil  {
		fmt.Println("error in reading argument")
		return
	}
	conten := strings.TrimSpace(string(file))
	line := strings.Split(conten, "\n")
	fmt.Println(line)
}