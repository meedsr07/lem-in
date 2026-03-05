package antmoving

import "fmt"
type Ant struct {
	ID  int
	Pos int
	Path []string
}

func MoveAntsCorrect(paths []PathInfo) {
	var ants []Ant
	nextAntID := 1
	totalAnts := 0
	for _, p := range paths {
		totalAnts += p.Ants
	}

	antsInEnd := 0
	turn := 1

	// مصفوفة لتخزين كم نمل دخل كل مسار
	entered := make([]int, len(paths))

	for antsInEnd < totalAnts {
		line := ""
		// أولاً، نضيف نمل جديد لكل مسار إذا ممكن
		for i, p := range paths {
			if entered[i] < p.Ants {
				ants = append(ants, Ant{ID: nextAntID, Pos: 0, Path: p.Path})
				nextAntID++
				entered[i]++
			}
		}

		// الآن نحرك كل النمل
		for i := range ants {
			if ants[i].Pos < len(ants[i].Path)-1 {
				ants[i].Pos++
				room := ants[i].Path[ants[i].Pos]
				line += fmt.Sprintf("L%d-%s ", ants[i].ID, room)
				if ants[i].Pos == len(ants[i].Path)-1 {
					antsInEnd++
				}
			}
		}

		if line != "" {
			fmt.Println(line)
			turn++
		}
	}
}