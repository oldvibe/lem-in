package ants

import (
	"sort"
	"strings"
)

func GetShortestPaths(Paths [][]string) [][]string {
	sort.Slice(Paths, func(i, j int) bool {
		return len(Paths[i]) < len(Paths[j])
	})

	var ShortestPaths [][]string
	PathsSet := make(map[string]bool)

	for _, path := range Paths {
		isUnique := true
		pathKey := strings.Join(path, ",")

		if PathsSet[pathKey] {
			continue
		}

		for _, shortPath := range ShortestPaths {
			if MatchAnyRoom(path, shortPath) {
				isUnique = false
				break
			}
		}

		if isUnique {
			ShortestPaths = append(ShortestPaths, path)
			PathsSet[pathKey] = true
		}

	}

	return ShortestPaths
}
