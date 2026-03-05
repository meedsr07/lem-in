package antmoving


func SelctionPaths(allPath [][]string) [][]string {
	var PathsSelected [][]string
	for _, path := range allPath {
		var isIntersecting bool
		for _, selectedPath := range PathsSelected {
			if Intersection(path, selectedPath) {
				isIntersecting = true
				break
			}
		}		
		if !isIntersecting {
			PathsSelected = append(PathsSelected, path)
		}
	}
	return PathsSelected
}

func Intersection(path1 , path2 []string) bool {
	for i := 1 ; i < len(path1) -1 ; i++ {
		for j := 1 ; j < len(path2) -1 ; j++ {
			if path1[i] == path2[j] {
				return true
			}
		}
	}
	return false
} 