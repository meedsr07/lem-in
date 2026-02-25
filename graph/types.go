package graph

// Room represents a room with coordinates and special properties
type Room struct {
	Name    string
	X       int
	Y       int
	IsStart bool
	IsEnd   bool
}
