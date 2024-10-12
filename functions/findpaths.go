package ants

func FindPaths(farm *AntFarm) [][]string {
	return DFS(farm.Start, farm.End, []string{}, farm.Connections, [][]string{})
}

func DFS(Start, End string, Path []string, Graph map[string][]string, result [][]string) [][]string {
	Path = append(Path, Start)
	if Start == End {
		DestinationPath := make([]string, len(Path))
		copy(DestinationPath, Path)
		return append(result, DestinationPath)
	}
	for _, NextRoom := range Graph[Start] {
		if !Contain(Path, NextRoom) {
			result = DFS(NextRoom, End, Path, Graph, result)
		}
	}
	return result
}

func Contain(Path []string, NextRoom string) bool {
	for i := 0; i < len(Path); i++ {
		if Path[i] == NextRoom {
			return true
		}
	}
	return false
}