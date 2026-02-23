package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckStartandEnd(lines []string) bool {
	foundstart := false
	foundend := false

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if line == "" || !strings.HasPrefix(line, "#") {
			continue
		}

		if line == "##start" {
			if i+1 >= len(lines) {
				fmt.Println("Error: ##start without room")
				return false
			}

			part := strings.Fields(lines[i+1])
			if len(part) != 3 {
				fmt.Println("Error Invalid format")
				return false
			}

			if part[0] == "" || strings.HasPrefix(part[0], "L") {
				fmt.Println("Invalid room name")
				return false
			}

			_, errX := strconv.Atoi(part[1])
			_, errY := strconv.Atoi(part[2])
			if errX != nil || errY != nil {
				fmt.Println("Error invalid information")
				return false
			}

			foundstart = true
		}

		if line == "##end" {
			if i+1 >= len(lines) {
				fmt.Println("Error: ##end without room")
				return false
			}

			part := strings.Fields(lines[i+1])
			if len(part) != 3 {
				fmt.Println("Error Invalid format")
				return false
			}

			if part[0] == "" || strings.HasPrefix(part[0], "L") {
				fmt.Println("Invalid room name")
				return false
			}

			_, errX := strconv.Atoi(part[1])
			_, errY := strconv.Atoi(part[2])
			if errX != nil || errY != nil {
				fmt.Println("Error invalid information")
				return false
			}

			foundend = true
		}
	}

	if !foundstart {
		fmt.Println("start room is not exist")
		return false
	}

	if !foundend {
		fmt.Println("end room is not exist")
		return false
	}

	return true
}


