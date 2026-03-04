package bfs


func BFSAllPaths(graph map[string][]string, start, end string) [][]string {
	queue := [][]string{{start}}
	var allPaths [][]string

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]

		if last == end {
			allPaths = append(allPaths, path)
		}

		for _, neighbor := range graph[last] {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return allPaths
}

func contains(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}
