package graph

import (
	"fmt"
	"strings"
)

func RoomLinksexist(lines []string, room []Room) bool {
	roomMap := make(map[string]bool)
	for _, r := range room {
		roomMap[r.Name] = true
	}
	for _, v := range lines {
		line := strings.TrimSpace(v)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "-") {
			part := strings.Split(line, "-")
			if len(part) != 2 {
				fmt.Println("Error: invalid link format")
				return false
			}
			rome1 := part[0]
			rome2 := part[1]
			if !roomMap[rome1] || !roomMap[rome2] {
				fmt.Println("link contains non-existing room")
				return false
			}
		}
	}
	return true
}
