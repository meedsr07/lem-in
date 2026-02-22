package main

import (
	"strconv"
	"strings"
)

type Room struct {
	Name    string
	X       int
	Y       int
	IsStart bool
	IsEnd   bool
}

func GetRoom(lines []string, startindex int) []Room {
	var res []Room
	isstart := false
	isend := false
	for i := startindex; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == ""  {
			continue
		}
		if line == "##start" {
			isstart = true
		}
		if line == "##end" {
			isend = true
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

		roomData := Room{
			Name: part[0],
			X:    x,
			Y:    y,
			IsStart: isstart,
			IsEnd: isend,
		}
		res = append(res, roomData)
		isstart = false
		isend = false

	}
	return res
}
