package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidForma(lines []string) bool {
	antsChecked := false

	for _, v := range lines {
		line := strings.TrimSpace(v)

		if line == "" {
			continue
		}

		
		if strings.HasPrefix(line, "#") {
			continue
		}

		// ants (first valid line)
		if !antsChecked {
			_, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("invalid ants")
				return false
			}
			antsChecked = true
			continue
		}

		// room (name x y)
		fields := strings.Fields(line)
		if len(fields) == 3 {
			if _, err1 := strconv.Atoi(fields[1]); err1 == nil {
				if _, err2 := strconv.Atoi(fields[2]); err2 == nil {
					continue
				}
			}
		}

		// link (room-room)
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				continue
			}
		}

		// anything else = error
		fmt.Println("invalid file format")
		return false
	}

	return true
}
