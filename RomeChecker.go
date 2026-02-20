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
func GetRome(lines []string , startindex int) ([]Rome , int) {
	var romes []Rome
for i := startindex ; i < len(lines) ; i++ {
	line := strings.TrimSpace(lines[i])


}
}


func CheckStartEnd(lines []string) bool {
	foundStart := false
	foundEnd := false

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if line == "##start" {
			if foundStart {
				return false 
			}
			if i+1 >= len(lines) {
				return false 
			}

			next := strings.TrimSpace(lines[i+1])
			fields := strings.Fields(next)
			if len(fields) != 3 {
				return false 
			}

			foundStart = true
		}

		if line == "##end" {
			if foundEnd {
				return false 
			}
			if i+1 >= len(lines) {
				return false
			}

			next := strings.TrimSpace(lines[i+1])
			fields := strings.Fields(next)
			if len(fields) != 3 {
				return false
			}

			foundEnd = true
		}
	}

	if !foundStart || !foundEnd {
		return false
	}

	return true
}