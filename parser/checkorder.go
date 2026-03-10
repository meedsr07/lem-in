package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckOrder(lines []string) bool {
	state := 1

	for _, v := range lines {
		line := strings.TrimSpace(v)

		if line == "" {
			continue
		}

		// skip comments
		if strings.HasPrefix(line, "#") {
			continue
		}

		// ants
		if state == 1 {
			_, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("invalid ants")
				return false
			}
			state = 2

			// rooms
		} else if state == 2 {
			if strings.Contains(line, "-") {
				state = 3
			} else {
				fields := strings.Fields(line)
				if len(fields) != 3 {
					fmt.Println("invalid room")
					return false
				}
			}

			// links
		} else if state == 3 {
			if !strings.Contains(line, "-") {
				fmt.Println("room after links not allowed")
				return false
			}
		}
	}

	return true
}
