package ants

func DFS(Start, End string, Path []string, Graph map[string][]string, result [][]string) [][]string {
	Path = append(Path, Start)
	if Start == End {
		DestinationPath := make([]string, len(Path))
		copy(DestinationPath, Path)
		return append(result, DestinationPath)
	}
	for _, NextRoom := range Graph[Start] {
		if !Contains(Path, NextRoom) {
			result = DFS(NextRoom, End, Path, Graph, result)
		}
	}
	return result
}