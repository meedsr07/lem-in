package parser

import (
	"fmt"
	"lem-in/bfs"
	"lem-in/graph"
	"os"
	"strings"
)

func Validation(arg string) {

	// reading file
	file, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println("error in reading file")
		return
	}
	// if file is empty
	if len(file) == 0 {
		fmt.Println("error : file is empty")
		return
	}
	// trim whit space in first and laste
	conten := strings.TrimSpace(string(file))
	// split file conten by \n
	lines := strings.Split(conten, "\n")

	antNbr, lineIndex := graph.GetAnts(lines)
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
	room := graph.GetRoom(lines, lineIndex)

	if !CheckDuplicateRooms(room) {
		return
	}
	if !RoomLinksexist(lines, room) {
		return
	}
	graphResult := graph.BulidGraph(lines, room)
	start, end := graph.GetStartandEnd(room)
	allpaht := bfs.BFSAllPaths(graphResult, start, end)
	if len(allpaht) == 0 {
		fmt.Println("ERROR: no path found")
		return
	}

	bestPaths := bfs.Filterpaths(allpaht)
	fmt.Println("best paths:", bestPaths)
	fmt.Println("Number of ants:", antNbr, "line index:", lineIndex, room)
	fmt.Println(graphResult)
	fmt.Println(start , end)
}
