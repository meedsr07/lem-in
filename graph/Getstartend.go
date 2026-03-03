package graph

func GetStartandEnd(room []Room) (string, string) {
	start := ""
	end := ""
	for _, r := range room {
		if r.IsStart {
			start = r.Name
		}
		if r.IsEnd {
			end = r.Name
		}
	}
	return start, end
}
