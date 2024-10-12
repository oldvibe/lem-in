package ants

import "fmt"

func validateFarm(farm *AntFarm) error {
	var startCount, endCount int
	for _, room := range farm.Rooms {
		if room.Type == "start" {
			startCount++
		} else if room.Type == "end" {
			endCount++
		}
	}
	if startCount != 1 {
		return fmt.Errorf("invalid number of start rooms: %d", startCount)
	}
	if endCount != 1 {
		return fmt.Errorf("invalid number of end rooms: %d", endCount)
	}
	return nil
}
