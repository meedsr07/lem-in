package antmoving

type PathInfo struct {
	Path []string 
	Ants int      
}

func DistributeAnts(paths [][]string, totalAnts int) []PathInfo {
	n := len(paths)
	result := make([]PathInfo, n)

	lengths := make([]int, n)
	for i, p := range paths {
		lengths[i] = len(p) - 1 
	}


	antsRemaining := totalAnts

	for antsRemaining > 0 {
		minIndex := 0
		minTurns := lengths[0] + result[0].Ants
		for i := 1; i < n; i++ {
			turns := lengths[i] + result[i].Ants
			if turns < minTurns {
				minTurns = turns
				minIndex = i
			}
		}
		result[minIndex].Path = paths[minIndex]
		result[minIndex].Ants++
		antsRemaining--
	}

	return result
}