package parser

import (
	"lem-in/graph"
	"strconv"
	"strings"
)

func GetRoom(lines []string, startindex int) []graph.Room {
	var res []graph.Room
	isstart := false
	isend := false
	for i := startindex; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		if line == "##start" {
			isstart = true
			continue
		}
		if line == "##end" {
			isend = true
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "-") {
			break
		}
		part := strings.Fields(line)
		if len(part) != 3 {
			continue
		}
		x, _ := strconv.Atoi(part[1])
		y, _ := strconv.Atoi(part[2])

		roomData := graph.Room{
			Name:    part[0],
			X:       x,
			Y:       y,
			IsStart: isstart,
			IsEnd:   isend,
		}
		res = append(res, roomData)
		isstart = false
		isend = false

	}
	return res
}
