package parser

import (
	"fmt"
	"lem-in/bfs"
	"lem-in/graph"
	"lem-in/antmoving"
	"os"
	"strings"
)

func Validation(arg string) {

	// -------- READ FILE --------
	file, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	if len(file) == 0 {
		fmt.Println("ERROR")
		return
	}

	content := strings.TrimSpace(string(file))
	lines := strings.Split(content, "\n")

	// -------- PARSE ANTS --------
	antNbr, lineIndex := graph.GetAnts(lines)
	if lineIndex == -1 || antNbr <= 0 {
		fmt.Println("ERROR")
		return
	}

	// -------- VALIDATION --------
	if !CheckStartandEnd(lines) ||
		!LinkChecker(lines) {
		fmt.Println("ERROR")
		return
	}

	rooms := graph.GetRoom(lines, lineIndex)

	if !CheckDuplicateRooms(rooms) ||
		!RoomLinksexist(lines, rooms) {
		fmt.Println("ERROR")
		return
	}

	// -------- BUILD GRAPH --------
	graphResult := graph.BulidGraph(lines, rooms)
	start, end := graph.GetStartandEnd(rooms)

	// -------- FIND ALL PATHS --------
	allPaths := bfs.BFSAllPaths(graphResult, start, end)
	if len(allPaths) == 0 {
		fmt.Println("ERROR")
		return
	}
	selectedPath := antmoving.SelctionPaths(allPaths)
	Pathinfo := antmoving.DistributeAnts(selectedPath, antNbr)
	fmt.Println(content)
	fmt.Println()
	antmoving.MoveAntsCorrect(Pathinfo)

}