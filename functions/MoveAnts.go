package ants

import (
	"fmt"
	"strings"
)

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
