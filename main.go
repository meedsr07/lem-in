package main

import (
	"fmt"
	"os"
	"strconv"
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
	if len(file) == 0 {
		fmt.Println("error empty file")
		return
	}
	conten := strings.TrimSpace(string(file))
	lines := strings.Split(conten, "\n")

	antNbr := GetAnts(lines)
	if antNbr == 0 {
		return
	}
	fmt.Printf("Number of ants: %d\n", antNbr)

}

func GetAnts(lines []string) int {
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" || strings.HasPrefix(line, "#") {
            continue 
        }

        fields := strings.Fields(line)
        if len(fields) != 1 {
            fmt.Println("ERROR: invalid data format")
            return 0
        }

        antNbr, err := strconv.Atoi(fields[0])
        if err != nil || antNbr <= 0 {
            fmt.Println("ERROR: invalid data format")
            return 0
        }

        return antNbr 
    }

    fmt.Println("ERROR: invalid data format") 
    return 0
}
