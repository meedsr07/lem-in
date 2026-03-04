package bfs

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