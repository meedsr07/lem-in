package main

import (
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X    int
	Y    int
}

func GetRoom(lines []string, startindex int) []Room {
	var res []Room
	for i := startindex; i < len(lines); i++ {
		line  := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line,"#") {
			continue
		}
		if strings.Contains(line , "-") {
			break
		}
		part := strings.Fields(line)
		if len(part) != 3 {
			continue
		}
		x , _ := strconv.Atoi(part[1])
		y , _ := strconv.Atoi(part[2])

		roomData := Room {
			Name: part[0],
			X: x,
			Y: y,
		}
		res = append(res, roomData)

	}
	return res
}
