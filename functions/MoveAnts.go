package ants

import (
	"fmt"
	"strings"
)

// func SendAnts(Paths [][]string, Ants int) {
// 	Ant := 1
// 	PathsLength := len(Paths)
// 	PhatsIndexLength := make([][]int, PathsLength)
// 	for i := 0; i < PathsLength && Ant <= Ants; i++ {
// 		if i != PathsLength-1 {
// 			if len(PhatsIndexLength[i])+len(Paths[i]) > len(PhatsIndexLength[i+1])+len(Paths[i+1]) {
// 				continue
// 			}
// 		}
// 		if i != 0 {
// 			if len(PhatsIndexLength[i])+len(Paths[i]) >= len(PhatsIndexLength[i-1])+len(Paths[i-1]) {
// 				i = -1
// 				continue
// 			}
// 		}
// 		PhatsIndexLength[i] = append(PhatsIndexLength[i], Ant)
// 		Ant++
// 		i--
// 	}
// 	curr := 0
// 	Position := map[int]int{}
// 	Ant = 1
// 	for i := 0; i < len(PhatsIndexLength); i++ {
// 		for j := 0; j <= curr && j < len(PhatsIndexLength[i]); j++ {
// 			Position[PhatsIndexLength[i][j]]++
// 			if Position[PhatsIndexLength[i][j]] < len(Paths[i]) {
// 				if Position[PhatsIndexLength[i][j]] == len(Paths[i])-1 {
// 					Ant++
// 				}
// 				fmt.Print("L", PhatsIndexLength[i][j], "-", Paths[i][Position[PhatsIndexLength[i][j]]])
// 				if i != len(PhatsIndexLength)-1 {
// 					fmt.Print(" ")
// 				}
// 			}
// 		}
// 		curr++
// 		if i == len(PhatsIndexLength)-1 && Ant <= Ants-1 {
// 			i = -1
// 			fmt.Println()
// 		}
// 	}
// 	fmt.Println()
// }

func MoveAntsOld(farm *AntFarm, paths [][]string) {
	antPositions := make(map[int]int)
	for ant := 1; ant <= farm.NumAnts; ant++ {
		antPositions[ant] = -1
	}

	for {
		moved := false
		moves := make(map[int]string)

		for ant := 1; ant <= farm.NumAnts; ant++ {
			if antPositions[ant] < len(paths[0])-1 {
				antPositions[ant]++
				moves[ant] = paths[0][antPositions[ant]]
				moved = true
			}
		}
		if len(moves) == 0 {
			break
		}

		if !moved {
			break
		}
		fmt.Println("\033[m" + "=== MOVEMENT ===" + "\033[0m")

		var movesSlice []string
		for ant, room := range moves {
			movesSlice = append(movesSlice, fmt.Sprintf("\033[33m"+"L%d<=====>%s "+"\033[0m", ant, room))
		}
		fmt.Println(strings.Join(movesSlice, " "))
	}
}
