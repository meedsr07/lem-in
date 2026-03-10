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
	
	file, err := os.ReadFile(arg)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	if len(file) == 0 {
		fmt.Println("ERROR empty file")
		return
	}

	content := strings.TrimSpace(string(file))
	lines := strings.Split(content, "\n")
	if !ValidForma(lines){
		return
	}
	if !CheckOrder(lines) {
		return
	}
	// -------- PARSE ANTS --------
	antNbr, lineIndex := graph.GetAnts(lines)
	if lineIndex == -1 || antNbr <= 0 {
		return
	}

	// -------- VALIDATION --------
	if !CheckStartandEnd(lines) || 	!LinkChecker(lines) {
		return
	}

	rooms := graph.GetRoom(lines, lineIndex)

	if !CheckDuplicateRooms(rooms) ||!RoomLinksexist(lines, rooms) {
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
	selectedPath := antmoving.SelectBestPaths(allPaths,antNbr)
	Pathinfo := antmoving.DistributeAnts(selectedPath, antNbr)
	fmt.Println(content)
	fmt.Println()
	antmoving.MoveAntsCorrect(Pathinfo)

}