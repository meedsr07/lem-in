package main

import (
	"fmt"
	"os"
	"strings"
)

func Validation(arg string) {
	file, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println("error in reading file")
		return
	}
	if len(file) == 0 {
		fmt.Println("error : file is empty")
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

	if !LinkChecker(lines) {
		return
	}
	room := GetRoom(lines, lineIndex)

	if !CheckDuplicateRooms(room) {
		return
	}

	fmt.Println("Number of ants:", antNbr, "line index:", lineIndex, room)
}
