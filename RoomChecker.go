package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Rome struct {
	Name string
	X    int
	Y    int
}

func GetRomes(lines []string, startindx int) []Rome {
	for i := startindx; i < len(lines); i++ {
	}
}

func CheckStartandEnd(lines []string) bool {
	foundstart := false
	founend := false
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if line == "##start" {
			if i < len(line) {
				part := strings.Fields(string(line[i+1]))
				if len(part) != 3 {
					fmt.Println("Error Invalid forma")
					return false
				}
				X := part[1]
				Y := part[2]
				_, errX := strconv.Atoi(X)
				_, errY := strconv.Atoi(Y)
				if errX != nil || errY != nil {
					fmt.Println("Error invalid informa")
					return false
				}
			}
			foundstart = true
		}
		if line == "##end" {
			if i < len(line) {
				part := strings.Fields(string(line[i+1]))
				if len(part) != 3 {
					fmt.Println("Error Invalid forma")
					return false
				}
				X := part[1]
				Y := part[2]
				_, errX := strconv.Atoi(X)
				_, errY := strconv.Atoi(Y)
				if errX != nil || errY != nil {
					fmt.Println("Error invalid informa")
					return false
				}
			}
			founend = true
		}

		if !foundstart {
			fmt.Println("start rome  is not exist")
			return false
		}
		if !founend {
			fmt.Println("end rome is not exist")
			return false
		}

	}
	return true
}
