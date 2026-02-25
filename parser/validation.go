package parser

import (
	"fmt"
	"lem-in/graph"
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
	if !graph.CheckStartandEnd(lines) {
		return
	}

	if !graph.LinkChecker(lines) {
		return
	}
	room := GetRoom(lines, lineIndex)

	if !CheckDuplicateRooms(room) {
		return
	}
	if !graph.RoomLinksexist(lines, room) {
		return
	}
	graphResult := graph.BulidGraph(lines, room)

	fmt.Println("Number of ants:", antNbr, "line index:", lineIndex, room)
	fmt.Println(graphResult)
}
