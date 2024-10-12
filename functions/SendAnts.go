package ants

import (
	"fmt"
)

func SendAnts(Paths [][]string, Ants int) {
	Ant := 1
	PathsLength := len(Paths)
	PhatsIndexLength := make([][]int, PathsLength)
	for i := 0; i < PathsLength && Ant <= Ants; i++ {
		if i != PathsLength-1 {
			if len(PhatsIndexLength[i])+len(Paths[i]) > len(PhatsIndexLength[i+1])+len(Paths[i+1]) {
				continue
			}
		}
		if i != 0 {
			if len(PhatsIndexLength[i])+len(Paths[i]) >= len(PhatsIndexLength[i-1])+len(Paths[i-1]) {
				i = -1
				continue
			}
		}
		PhatsIndexLength[i] = append(PhatsIndexLength[i], Ant)
		Ant++
		i--
	}
	curr := 0
	Position := map[int]int{}
	Ant = 1
	for i := 0; i < len(PhatsIndexLength); i++ {
		for j := 0; j <= curr && j < len(PhatsIndexLength[i]); j++ {
			Position[PhatsIndexLength[i][j]]++
			if Position[PhatsIndexLength[i][j]] < len(Paths[i]) {
				if Position[PhatsIndexLength[i][j]] == len(Paths[i])-1 {
					Ant++
				}
				fmt.Print("L", PhatsIndexLength[i][j], "-", Paths[i][Position[PhatsIndexLength[i][j]]])
				if i != len(PhatsIndexLength) {
					fmt.Print(" ")
				}
			}
		}
		curr++
		if i == len(PhatsIndexLength)-1 && Ant <= Ants{
			i = -1
			fmt.Println()
		}
	}
	fmt.Println()
}