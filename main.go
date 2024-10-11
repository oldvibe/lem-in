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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X, Y int
	Type string // "start", "end", or ""
}

type AntFarm struct {
	NumAnts     int
	Rooms       map[string]*Room
	Connections map[string][]string
	Start       string
	End         string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	farm, err := parseInput(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: invalid data format, %v\n", err)
		return
	}

	// Print the farm structure
	printFarm(farm)

	// Find paths and move ants
	paths := findPaths(farm)
	if len(paths) == 0 {
		fmt.Println("ERROR: no valid path found")
		return
	}

	moveAnts(farm, paths)
}

func parseInput(filename string) (*AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	farm := &AntFarm{
		Rooms:       make(map[string]*Room),
		Connections: make(map[string][]string),
	}

	scanner := bufio.NewScanner(file)
	
	// Parse number of ants
	if !scanner.Scan() {
		return nil, fmt.Errorf("empty file")
	}
	if farm.NumAnts, err = strconv.Atoi(scanner.Text()); err != nil || farm.NumAnts <= 0 {
		return nil, fmt.Errorf("invalid number of ants")
	}

	// Parse rooms and connections
	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" || line == "##end" {
			if err := parseSpecialRoom(farm, scanner, line); err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "-") {
			if err := parseConnection(farm, line); err != nil {
				return nil, err
			}
		} else if !strings.HasPrefix(line, "#") {
			if err := parseRoom(farm, line, ""); err != nil {
				return nil, err
			}
		}
	}

	if err := validateFarm(farm); err != nil {
		return nil, err
	}

	return farm, nil
}

func parseSpecialRoom(farm *AntFarm, scanner *bufio.Scanner, roomType string) error {
	if !scanner.Scan() {
		return fmt.Errorf("missing room after %s", roomType)
	}
	roomType = strings.TrimPrefix(roomType, "##")
	if err := parseRoom(farm, scanner.Text(), roomType); err != nil {
		return err
	}
	if roomType == "start" {
		farm.Start = strings.Fields(scanner.Text())[0]
	} else {
		farm.End = strings.Fields(scanner.Text())[0]
	}
	return nil
}

func parseRoom(farm *AntFarm, line, roomType string) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("invalid room format: %s", line)
	}
	name, x, y := parts[0], parts[1], parts[2]
	if name[0] == 'L' || name[0] == '#' {
		return fmt.Errorf("invalid room name: %s", name)
	}
	xCoord, err := strconv.Atoi(x)
	if err != nil {
		return fmt.Errorf("invalid x coordinate: %s", x)
	}
	yCoord, err := strconv.Atoi(y)
	if err != nil {
		return fmt.Errorf("invalid y coordinate: %s", y)
	}
	if _, exists := farm.Rooms[name]; exists {
		return fmt.Errorf("duplicate room: %s", name)
	}
	farm.Rooms[name] = &Room{Name: name, X: xCoord, Y: yCoord, Type: roomType}
	return nil
}

func parseConnection(farm *AntFarm, line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid connection format: %s", line)
	}
	room1, room2 := parts[0], parts[1]
	if _, exists := farm.Rooms[room1]; !exists {
		return fmt.Errorf("unknown room in connection: %s", room1)
	}
	if _, exists := farm.Rooms[room2]; !exists {
		return fmt.Errorf("unknown room in connection: %s", room2)
	}
	farm.Connections[room1] = append(farm.Connections[room1], room2)
	farm.Connections[room2] = append(farm.Connections[room2], room1)
	return nil
}

func validateFarm(farm *AntFarm) error {
	if farm.Start == "" {
		return fmt.Errorf("no start room found")
	}
	if farm.End == "" {
		return fmt.Errorf("no end room found")
	}
	return nil
}

func printFarm(farm *AntFarm) {
	fmt.Println(farm.NumAnts)
	for _, room := range farm.Rooms {
		if room.Type == "start" {
			fmt.Println("##start")
		} else if room.Type == "end" {
			fmt.Println("##end")
		}
		fmt.Printf("%s %d %d\n", room.Name, room.X, room.Y)
	}
	for room, connections := range farm.Connections {
		for _, conn := range connections {
			if room < conn {
				fmt.Printf("%s-%s\n", room, conn)
			}
		}
	}
	fmt.Println()
}

func findPaths(farm *AntFarm) [][]string {
	// Implement BFS to find the shortest path(s)
	visited := make(map[string]bool)
	queue := [][]string{{farm.Start}}
	var paths [][]string

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		room := path[len(path)-1]

		if room == farm.End {
			paths = append(paths, path)
			continue
		}

		if !visited[room] {
			visited[room] = true
			for _, nextRoom := range farm.Connections[room] {
				if !visited[nextRoom] {
					newPath := make([]string, len(path))
					copy(newPath, path)
					queue = append(queue, append(newPath, nextRoom))
				}
			}
		}
	}

	return paths
}

func moveAnts(farm *AntFarm, paths [][]string) {
	antPositions := make(map[int]int) // ant number -> position in path
	for ant := 1; ant <= farm.NumAnts; ant++ {
		antPositions[ant] = -1 // Start before the path
	}

	for {
		moved := false
		moves := make(map[int]string) // ant number -> room moved to

		for ant := 1; ant <= farm.NumAnts; ant++ {
			if antPositions[ant] < len(paths[0])-1 {
				antPositions[ant]++
				moves[ant] = paths[0][antPositions[ant]]
				moved = true
			}
		}

		if !moved {
			break
		}

		var movesSlice []string
		for ant, room := range moves {
			movesSlice = append(movesSlice, fmt.Sprintf("L%d-%s", ant, room))
		}
		fmt.Println(strings.Join(movesSlice, " "))
	}
}