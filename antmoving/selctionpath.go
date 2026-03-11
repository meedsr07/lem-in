package antmoving


func SelectBestPaths(allPaths [][]string, ants int) [][]string {
	var best [][]string
	bestTurns := 1<<31 - 1

	for i := 0; i < len(allPaths); i++ {

		var current [][]string
		current = append(current, allPaths[i])

		for j := i + 1; j < len(allPaths); j++ {

			intersect := false

			for _, p := range current {
				if Intersection(p, allPaths[j]) {
					intersect = true
					break
				}
			}

			if !intersect {
				current = append(current, allPaths[j])
			}
		}

		turns := CalculateTurns(current, ants)

		if turns < bestTurns {
			bestTurns = turns
			best = current
		}
	}

	return best
}

// CalculateTurns calculates the number of turns required to move all ants through the given paths.
func CalculateTurns(paths [][]string, ants int) int {
	lengths := make([]int, len(paths))

	for i := range paths {
		lengths[i] = len(paths[i]) - 2
	}

	antsOnPath := make([]int, len(paths))

	for a := 0; a < ants; a++ {

		best := 0
		for i := 1; i < len(paths); i++ {
			if lengths[i]+antsOnPath[i] < lengths[best]+antsOnPath[best] {
				best = i
			}
		}

		antsOnPath[best]++
	}

	maxTurns := 0
	for i := range paths {
	
		turns := lengths[i] + antsOnPath[i]

		if turns > maxTurns {
			maxTurns = turns
		}
	}

	return maxTurns
}

func Intersection(path1, path2 []string) bool {
	for i := 1; i < len(path1)-1; i++ {
		for j := 1; j < len(path2)-1; j++ {
			if path1[i] == path2[j] {
				return true
			}
		}
	}
	return false
}