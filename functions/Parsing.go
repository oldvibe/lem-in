package ants

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X, Y int
	Type string
}

type AntFarm struct {
	NumAnts     int
	Rooms       map[string]*Room
	Connections map[string][]string
	Start       string
	End         string
}

func ParseSpecialRoom(farm *AntFarm, scanner *bufio.Scanner, roomType string) error {
	if !scanner.Scan() {
		return fmt.Errorf("missing room after %s", roomType)
	}
	roomType = strings.TrimPrefix(roomType, "##")
	if err := ParseRoom(farm, scanner.Text(), roomType); err != nil {
		return err
	}
	if roomType == "start" {
		farm.Start = strings.Fields(scanner.Text())[0]
	} else {
		farm.End = strings.Fields(scanner.Text())[0]
	}
	return nil
}

func ParseRoom(farm *AntFarm, line, roomType string) error {
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

func ParseConnection(farm *AntFarm, line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid link format: %s", line)
	}
	room1 := parts[0]
	room2 := parts[1]
	if _, exists := farm.Rooms[room1]; !exists {
		return fmt.Errorf("unknown room in link: %s", room1)
	}
	if _, exists := farm.Rooms[room2]; !exists {
		return fmt.Errorf("unknown room in link: %s", room2)
	}
	farm.Connections[room1] = append(farm.Connections[room1], room2)
	farm.Connections[room2] = append(farm.Connections[room2], room1)
	return nil
}
