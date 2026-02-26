package bfs

func FindPath(graph map[string][]string, start string, end string) []string {
	queue := []string{start}
	visted := make(map[string]bool)
	visted[start] = true
	// for remembers the room you came from for each room
	parant := make(map[string]string)
	//     as long as there is at least one room in the queue, keep going
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range graph[current] {
			if !visted[neighbor] {
				visted[neighbor] = true
				parant[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	if !visted[end] {
		return nil
	}

	path := []string{}
	current := end
	for current != "" {
		path = append([]string{current}, path...)
		current = parant[current]
	}

	return path

}
