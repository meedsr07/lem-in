package parser

import "lem-in/graph"


func GetStartandEnd(room []graph.Room) (string , string) {
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
return  start , end
}