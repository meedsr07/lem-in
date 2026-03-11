package antmoving

import "fmt"
type Ant struct {
	ID  int
	Pos int
	Path []string
}

func MoveAntsCorrect(paths []PathInfo) {
	// Create a slice to hold the ants
	var ants []Ant
// Calculate the total number of ants and initialize the next ant ID
	nextAntID := 1
	// Calculate the total number of ants
	totalAnts := 0
	// Calculate the total number of ants by summing the Ants field of each PathInfo
	for _, p := range paths {
		totalAnts += p.Ants
	}
	// Initialize the number of ants that have reached the end and the turn counter
	antsInEnd := 0
	// Loop until all ants have reached the end
	turn := 1
	// Create a slice to track how many ants have entered each path
	entered := make([]int, len(paths))
	// Loop until all ants have reached the end
	for antsInEnd < totalAnts {
		// Create a string to hold the output for this turn
		line := ""
	// Loop through each path and add new ants to the path if there are still ants to enter
		for i, p := range paths {
			// Check if we can add another ant to this path (if we haven't already added all ants for this path)
			if entered[i] < p.Ants {
				// Add a new ant to the path with the next available ID and starting position
				ants = append(ants, Ant{ID: nextAntID, Pos: 0, Path: p.Path})
				// Increment the next ant ID and the count of ants that have entered this path
				nextAntID++
				// Increment the count of ants that have entered this path
				entered[i]++
				// Add the move of the new ant to the output line
			}
		}

	// Move each ant one step forward if it hasn't reached the end of its path
		for i := range ants {
			// Check if the ant has not reached the end of its path
			if ants[i].Pos < len(ants[i].Path)-1 {
				// Move the ant one step forward
				ants[i].Pos++
				// Add the move of the ant to the output line
				room := ants[i].Path[ants[i].Pos]
				// Add the move of the ant to the output line in the format "L{antID}-{room}"
				line += fmt.Sprintf("L%d-%s ", ants[i].ID, room)
				// Check if the ant has reached the end of its path after moving
				if ants[i].Pos == len(ants[i].Path)-1 {
					// Increment the count of ants that have reached the end
					antsInEnd++
					// Remove the ant from the slice of ants (optional, can be left in since we check position)
				}
			}
		}
		// Print the output line for this turn if it's not empty

		if line != "" {
			// Trim the trailing space from the line
			fmt.Println(line)
			// Increment the turn counter
			turn++
		}
	}
}