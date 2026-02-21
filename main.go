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

	antNbr, lineIndex := GetAnts(lines)
	if lineIndex == -1 {
		return
	}
	if lineIndex == len(lines)-1 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	if antNbr == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	if !CheckStartandEnd(lines) {
		return
	}
	room := GetRoom(lines , lineIndex)
	
	fmt.Println("Number of ants:",antNbr , "line index:", lineIndex , room)

}

func GetAnts(lines []string) (int , int) {
    for i , line := range lines {
        line = strings.TrimSpace(line)
        if line == "" || strings.HasPrefix(line, "#") {
            continue 
        }

        fields := strings.Fields(line)
        if len(fields) != 1 {
            fmt.Println("ERROR: invalid data format")
            return 0 , -1
        }

        antNbr, err := strconv.Atoi(fields[0])
        if err != nil || antNbr <= 0 {
            fmt.Println("ERROR: invalid data format")
            return 0 , -1
        }

        return antNbr , i
    }

    fmt.Println("ERROR: invalid data format") 
    return 0 , -1
}

