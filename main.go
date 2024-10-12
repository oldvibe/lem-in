package main

import (
	"fmt"
	"os"

	ants "ants/functions"
)

type (
	AntFarm = ants.AntFarm
	Room    = ants.Room
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	farm, err := ants.ParseInput(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: invalid data format, %v\n", err)
		return
	}
	ants.PrintFarm(farm)

	paths := ants.FindPaths(farm)
	Unique := ants.GetShortestPaths(paths)
	for _, path := range Unique {
		fmt.Print( "\033[35m")
		fmt.Println(len(path))
		fmt.Print( "\033[0m")
		fmt.Print( "\033[32m")
		fmt.Println(path)
		fmt.Println( "\033[0m")
	}
	if len(Unique) == 0 {
		fmt.Println("ERROR: no valid path found")
		return
	}
	ants.PrintResult(Unique, farm.NumAnts)
	//ants.SendAnts(Unique, farm.NumAnts)
}
