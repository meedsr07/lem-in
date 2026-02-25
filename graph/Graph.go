package graph

import (
	"strings"
)

func BulidGraph(lines []string, rooms []Room) map[string][]string {
	Graph := make(map[string][]string)
	roomMap := make(map[string]bool)
	for _, r := range rooms {
		roomMap[r.Name] = true
	}
	for _, v := range lines {
		line := strings.TrimSpace(v)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "-") {
			part := strings.Split(line, "-")
			rome1 := part[0]
			rome2 := part[1]
			if !roomMap[rome1] || !roomMap[rome2] {
				continue
			}
			Graph[rome1] = append(Graph[rome1], rome2)
			Graph[rome2] = append(Graph[rome2], rome1)
		}
	}
	return Graph
}
