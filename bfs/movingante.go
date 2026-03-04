package bfs

import "fmt"

func Filterpaths(paths [][]string) [][]string {
	var bestPaths [][]string
	usedRooms := make(map[string]bool)

	for _, path := range paths {

		if !hasConflict(path, usedRooms) {
			bestPaths = append(bestPaths, path)

			// mark middle rooms as used
			for i := 1; i < len(path)-1; i++ {
				usedRooms[path[i]] = true
			}
		}
	}

	return bestPaths
}

func hasConflict(path []string, usedRooms map[string]bool) bool {
	for i := 1; i < len(path)-1; i++ {
		if usedRooms[path[i]] {
			return true
		}
	}
	return false
}

func DistributeAnts(paths [][]string, ants int) []int {
	counts := make([]int, len(paths))

	for ants > 0 {
		bestIndex := 0

		for i := 1; i < len(paths); i++ {
			if len(paths[i])+counts[i] < len(paths[bestIndex])+counts[bestIndex] {
				bestIndex = i
			}
		}

		counts[bestIndex]++
		ants--
	}

	return counts
}

func Simulate(paths [][]string, antCounts []int, totalAnts int) {
	type PathState struct {
		rooms []string
		ants  []int
	}

	var states []PathState
	antID := 1

	for i, p := range paths {
		states = append(states, PathState{
			rooms: p,
			ants:  make([]int, len(p)),
		})
		_ = i
	}

	finished := 0

	for finished < totalAnts {
		moves := ""

		for i := range states {
			path := &states[i]

			// نحرك من الأخير للبداية
			for j := len(path.rooms) - 1; j > 0; j-- {
				if path.ants[j-1] != 0 && path.ants[j] == 0 {
					path.ants[j] = path.ants[j-1]
					path.ants[j-1] = 0

					moves +=
						"L" + itoa(path.ants[j]) +
							"-" + path.rooms[j] + " "

					if j == len(path.rooms)-1 {
						finished++
					}
				}
			}

			// إدخال نملة جديدة
			if antCounts[i] > 0 && path.ants[1] == 0 {
				path.ants[1] = antID
				moves +=
					"L" + itoa(antID) +
						"-" + path.rooms[1] + " "

				antID++
				antCounts[i]--
			}
		}

		if moves != "" {
			println(moves)
		}
	}
}

func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}
