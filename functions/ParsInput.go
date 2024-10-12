package ants

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(filename string) (*AntFarm, error) {
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
			if err := ParseSpecialRoom(farm, scanner, line); err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "-") {
			if err := ParseConnection(farm, line); err != nil {
				return nil, err
			}
		} else if !strings.HasPrefix(line, "#") {
			if err := ParseRoom(farm, line, ""); err != nil {
				return nil, err
			}
		}
	}

	if err := validateFarm(farm); err != nil {
		return nil, err
	}

	return farm, nil
}
