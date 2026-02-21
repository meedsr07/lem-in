package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetAnts(lines []string) (int, int) {
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 1 {
			fmt.Println("ERROR: invalid data format")
			return 0, -1
		}

		antNbr, err := strconv.Atoi(fields[0])
		if err != nil || antNbr <= 0 {
			fmt.Println("ERROR: invalid data format")
			return 0, -1
		}

		return antNbr, i
	}

	fmt.Println("ERROR: invalid data format")
	return 0, -1
}
