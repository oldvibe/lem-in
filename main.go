// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type Graph struct {
// 	NumAnts int
// 	Rooms   map[string]*Room
// 	Connections  map[string][]string
// }

// type Room struct {
// 	Name string
// 	X, Y int
// 	Type string
// }

// func main() {
// 	farm, err := parseInput("test.txt")

// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Parse farm with %d ants\n  ", farm.NumAnts)
// }

// func parseInput(filename string) (*Graph, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	farm := &Graph{
// 		Rooms: make(map[string]*Room),
// 		Connections: make(map[string][]string),
// 	}

// 	scanner := bufio.NewScanner(file)

// 	if !scanner.Scan() {
// 		return nil, fmt.Errorf("empty file")
// 	}
// 	numAnts, err := strconv.Atoi(scanner.Text())
// 	if err != nil || numAnts <= 0 {
// 		return nil, fmt.Errorf("invalid number of ants: %s", scanner.Text())
// 	}
// 	farm.NumAnts = numAnts

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line == "##start" || line == "##end" {
// 			if !scanner.Scan() {
// 				return nil, fmt.Errorf("missing room after %s", line)
// 			}
// 			roomType := "start"
// 			if line == "##end" {
// 				roomType = "end"
// 			}
// 			if err := ParseRoom(farm, scanner.Text(), roomType); err != nil {
// 				return nil, err
// 			}
// 		} else if strings.Contains(line, "-") {

// 			if err := parseLinks(farm, line); err != nil {
// 				return nil, err
// 			}
// 			break
// 		} else if !strings.HasPrefix(line, "#") {
// 			if err := ParseRoom(farm, line, "normal"); err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.HasPrefix(line, "#") {
// 			continue
// 		}
// 		if err := parseLinks(farm, line); err != nil {
// 			return nil, err
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return nil, fmt.Errorf("error reading file: %v", err)
// 	}

// 	if err := validateFarm(farm); err != nil {
// 		return nil, err
// 	}

// 	return farm, nil

// }

// func parseLinks(farm *Graph, line string) error {
// 	parts := strings.Split(line, "-")
// 	if len(parts) != 2 {
// 		return fmt.Errorf("invalid link format: %s", line)
// 	}
// 	room1 := parts[0]
// 	room2 := parts[1]
// 	if _, exists := farm.Rooms[room1]; !exists {
// 		return fmt.Errorf("unknown room in link: %s", room1)
// 	}
// 	if _, exists := farm.Rooms[room2]; !exists {
// 		return fmt.Errorf("unknown room in link: %s", room2)
// 	}
// 	farm.Connections[room1] = append(farm.Connections[room1], room2)
// 	farm.Connections[room2] = append(farm.Connections[room2], room1)
// 	return nil
// }

// func validateFarm(farm *Graph) error {
// 	var startCount, endCount int
// 	for _, room := range farm.Rooms {
// 		if room.Type == "start" {
// 			startCount++
// 		} else if room.Type == "end" {
// 			endCount++
// 		}
// 	}
// 	if startCount != 1 {
// 		return fmt.Errorf("invalid number of start rooms: %d", startCount)
// 	}
// 	if endCount != 1 {
// 		return fmt.Errorf("invalid number of end rooms: %d", endCount)
// 	}
// 	return nil
// }

// func ParseRoom(farm *Graph, line, roomType string) error {
// 	parts := strings.Fields(line)
// 	if len(parts) != 3 {
// 		return fmt.Errorf("invalid room format: %s", line)
// 	}
// 	name := parts[0]
// 	if name[0] == 'L' || name[0] == '#' {
// 		return fmt.Errorf("invalid room name: %s", name)
// 	}
// 	x, err := strconv.Atoi(parts[1])
// 	if err != nil {
// 		return fmt.Errorf("invalid x coordinate: %s", parts[1])
// 	}
// 	y, err := strconv.Atoi(parts[2])
// 	if err != nil {
// 		return fmt.Errorf("invalid y coordinate: %s", parts[2])
// 	}
// 	if _, exists := farm.Rooms[name]; exists {
// 		return fmt.Errorf("duplicate room: %s", name)
// 	}

// 	farm.Rooms[name] = &Room{Name: name, X: x, Y: y, Type: roomType}
// 	for name, room := range farm.Rooms {
// 		fmt.Println(name, "==============>" , room)
// 	}
// 	return nil
// }

// func DFS(adjacencyList map[string][]string, currentVertex string, listOfVisited []string) []string {
// 	listOfVisited = append(listOfVisited, currentVertex)
// 	childs := adjacencyList[currentVertex]
// 	if len(childs) == 0 {
// 		return listOfVisited
// 	}

// 	for _, child := range childs {
// 		if !isContain(listOfVisited, child) {
// 			listOfVisited = DFS(adjacencyList, child, listOfVisited)
// 		}
// 	}

// 	return listOfVisited
// }

// func isContain(arr []string, target string) bool {
// 	if len(arr) == 0 {
// 		return false
// 	}

// 	for _, elem := range arr {
// 		if elem == target {
// 			return true
// 		}
// 	}

// 	return false
// }

// func TransformToAdjacencyList(listOfEdges []string) map[string][]string {
// 	var result map[string][]string = make(map[string][]string)

// 	if len(listOfEdges) == 0 {
// 		return result
// 	}

// 	for _, pairOfVertex := range listOfEdges {
// 		vertexes := strings.Split(pairOfVertex, "-")
// 		result[vertexes[0]] = append(result[vertexes[0]], vertexes[1])
// 		result[vertexes[1]] = append(result[vertexes[1]], vertexes[0])
// 	}

// 	return result
// }

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
	jt := 1
	// ants.FindPaths(farm)
	ants.PrintFarm(farm)
	fmt.Println(jt)
	jt++
	paths := ants.FindPaths(farm)

	Unique := GetShortestPaths(paths)
	for _, path := range Unique {
		fmt.Println(path)
	}
	fmt.Println(jt)
	jt++
	if len(Unique) == 0 {
		fmt.Println("ERROR: no valid path found")
		return
	}
	fmt.Println(Unique)
	jt++

	ants.PrintResult(Unique, farm.NumAnts)
	fmt.Println(jt)
	jt++

	ants.MoveAntsOld(farm, Unique)
}

func GetShortestPaths(Paths [][]string) (ShortestPaths [][]string) {
	PathsLenght := len(Paths)
	for i := 0; i < PathsLenght; i++ {
		AppendIt := true
		for j := 0; j < len(ShortestPaths); j++ {
			if MatchAnyRoom(Paths[i], ShortestPaths[j]) {
				AppendIt = false
				break
			}
		}
		if AppendIt {
			ShortestPaths = append(ShortestPaths, Paths[i])
		}
	}
	return
}

func MatchAnyRoom(RoomsOne, RoomsTwo []string) bool {
	for i := 1; i < len(RoomsOne)-1; i++ {
		for j := 1; j < len(RoomsTwo)-1; j++ {
			if RoomsOne[i] == RoomsTwo[j] && i != 0 && i != len(RoomsOne)-1 {
				return true
			}
		}
	}
	return false
}
