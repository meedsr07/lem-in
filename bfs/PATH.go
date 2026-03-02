package bfs


// BFS finds one shortest path
func bfsPath(graph map[string][]string, start, end string) []string {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current == end {
			return path
		}

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return nil
}

// findMultiplePaths finds multiple non-overlapping shortest paths
func FindMultiplePaths(graph map[string][]string, start, end string) [][]string {
	var allPaths [][]string
	// create a copy of the graph to modify
	graphCopy := make(map[string][]string)
	for k, v := range graph {
		graphCopy[k] = append([]string{}, v...)
	}

	for {
		path := bfsPath(graphCopy, start, end)
		if path == nil {
			break
		}
		allPaths = append(allPaths, path)

		// remove edges of this path from the graph to avoid overlap
		for i := 0; i < len(path)-1; i++ {
			from, to := path[i], path[i+1]
			newNeighbors := []string{}
			for _, n := range graphCopy[from] {
				if n != to {
					newNeighbors = append(newNeighbors, n)
				}
			}
			graphCopy[from] = newNeighbors
		}
	}

	return allPaths
}