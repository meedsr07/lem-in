package main

import "strings"

type Rome struct {
	Name string
	X    int
	Y    int
}

func GetRome(lines []string, lineIndex int) ([]Rome, int) {
	for i := lineIndex + 1 ; i < len(lines) ; i++ {
		line := strings.TrimSpace(lines[i]) 
	}
}
