package bfs


func BFSAllPaths(graph map[string][]string, start, end string) [][]string {
	// 	Initialize a queue with the starting path and a slice to hold all found paths
	queue := [][]string{{start}}
	// Create a slice to hold all found paths from start to end
	var allPaths [][]string
	// Loop until the queue is empty

	for len(queue) > 0 {
		// 	Dequeue the first path and get the last node in that path
		path := queue[0]
		/// Remove the first path from the queue
		queue = queue[1:]
		// Get the last node in the current path
		last := path[len(path)-1]
		// If the last node is the end node, add the current path to the list of all paths

		if last == end {
			allPaths = append(allPaths, path)
		}
		// Otherwise, enqueue new paths for each neighbor of the last node that is not already in the current path
		// Loop through the neighbors of the last node
		for _, neighbor := range graph[last] {
			// Check if the neighbor is not already in the current path to avoid cycles
			if !contains(path, neighbor) {
				// Create a new path by appending the neighbor to the current path and enqueue it
				newPath := append([]string{}, path...)
				// Append the neighbor to the new path
				newPath = append(newPath, neighbor)
				// Enqueue the new path
				queue = append(queue, newPath)

			}
		}
	}
	 // Return all found paths from start to end
	return allPaths
}
// Helper function to check if a node is already in the current path
func contains(path []string, node string) bool {
	// Loop through the nodes in the path to check if the node is already present
	for _, n := range path {
		// If the node is found in the path, return true
		if n == node {
			return true
		}
	}
	
	return false
}
