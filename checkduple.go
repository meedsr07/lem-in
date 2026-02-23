package main

import "fmt"

func CheckDuplicateRooms(rooms []Room) bool {
	for i := 0; i < len(rooms); i++ {
		for j := i + 1; j < len(rooms); j++ {

			if rooms[i].Name == rooms[j].Name {
				fmt.Println("error invalid dupl rome")
				return false
			}

			if rooms[i].X == rooms[j].X && rooms[i].Y == rooms[j].Y {
				fmt.Println("invalid cordini")
				return false
			}
		}
	}
	return true
}
