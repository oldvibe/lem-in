package ants

import (
	"fmt"
)

func PrintFarm(farm *AntFarm) {
	fmt.Println(farm.NumAnts)
	for _, room := range farm.Rooms {
		if room.Type == "end" {
			fmt.Println("##end")
		} else if room.Type == "start" {
			fmt.Println("##start")
		}
		fmt.Printf("%s %d %d\n", room.Name, room.X, room.Y)
	}
	for room, connections := range farm.Connections {
		for _, conn := range connections {
			if room < conn {
				fmt.Printf("%s==========>%s\n", room, conn)
			}
		}
	}
	fmt.Println()
}
