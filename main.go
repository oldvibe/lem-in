package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("")
		return
	}
	
	parseInput("test.txt")
}

type Room struct {
	Name string
	X, Y int
	Type string
}

type AntFarm struct {
	NumAnts int
	Rooms   map[string]*Room
	Links   map[string][]string
}

func parseInput(filename string) (*AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	farm := &AntFarm{
		Rooms: make(map[string]*Room),
		Links: make(map[string][]string),
	}

	scanner := bufio.NewScanner(file)

	// Parse number of ants
	if !scanner.Scan() {
		return nil, fmt.Errorf("empty file")
	}
	numAnts, err := strconv.Atoi(scanner.Text())
	if err != nil || numAnts <= 0 {
		return nil, fmt.Errorf("invalid number of ants: %s", scanner.Text())
	}
	farm.NumAnts = numAnts

	// Parse rooms
	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" || line == "##end" {
			if !scanner.Scan() {
				return nil, fmt.Errorf("missing room after %s", line)
			}
			roomType := "start"
			if line == "##end" {
				roomType = "end"
			}
			if err := parseRoom(farm, scanner.Text(), roomType); err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "-") {
			// Start of links section
			break
		} else if !strings.HasPrefix(line, "#") {
			if err := parseRoom(farm, line, "normal"); err != nil {
				return nil, err
			}
		}
	}

	// Parse links
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid link format: %s", line)
		}
		room1, room2 := parts[0], parts[1]
		if _, exists := farm.Rooms[room1]; !exists {
			return nil, fmt.Errorf("unknown room in link: %s", room1)
		}
		if _, exists := farm.Rooms[room2]; !exists {
			return nil, fmt.Errorf("unknown room in link: %s", room2)
		}
		farm.Links[room1] = append(farm.Links[room1], room2)
		farm.Links[room2] = append(farm.Links[room2], room1)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return farm, nil
}

func parseRoom(farm *AntFarm, line, roomType string) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("invalid room format: %s", line)
	}
	name := parts[0]
	if name[0] == 'L' || name[0] == '#' {
		return fmt.Errorf("invalid room name: %s", name)
	}
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid x coordinate: %s", parts[1])
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid y coordinate: %s", parts[2])
	}
	if _, exists := farm.Rooms[name]; exists {
		return fmt.Errorf("duplicate room: %s", name)
	}
	farm.Rooms[name] = &Room{Name: name, X: x, Y: y, Type: roomType}
	return nil
}
