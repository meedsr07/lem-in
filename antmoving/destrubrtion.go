package antmoving

type PathInfo struct {
	Path []string 
	Ants int      
}

func DistributeAnts(paths [][]string, totalAnts int) []PathInfo {
	// Initialize result with the same length as paths
	n := len(paths)
	// Create a slice to hold the distribution of ants for each path
	result := make([]PathInfo, n)
	// Calculate the length of each path (excluding the start and end)
	lengths := make([]int, n)
	// Calculate the length of each path (excluding the start and end)
	for i, p := range paths {
		lengths[i] = len(p) - 1 
	}
	// Distribute ants to paths based on their lengths

	antsRemaining := totalAnts
	// Loop until all ants are distributed
	for antsRemaining > 0 {
		// Find the path that will allow the next ant to reach the end in the fewest turns
		bestpath := 0
		// Calculate the number of turns for each path if we add one more ant to it
		minTurns := lengths[0] + result[0].Ants
		// Loop through the paths to find the one with the minimum turns
		for i := 1; i < n; i++ {
			// Calculate the number of turns if we add one more ant to this path
			turns := lengths[i] + result[i].Ants
			// 	Check if this path has fewer turns than the current minimum
			if turns < minTurns {
				// Update the minimum turns and the index of the path
				minTurns = turns
				// Update the index of the path with the minimum turns
				bestpath = i
				// If the turns are equal, we can choose the path with fewer ants currently assigned
			}
		}
		// Assign one ant to the path with the minimum turns
		result[bestpath].Path = paths[bestpath]
		// Increment the number of ants assigned to this path
		result[bestpath].Ants++
		// Decrement the number of remaining ants
		antsRemaining--
	}

	// Return the result with the distribution of ants for each path
	return result
}