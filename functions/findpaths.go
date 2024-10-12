package ants

func FindPaths(farm *AntFarm) [][]string {
	return DFS(farm.Start, farm.End, []string{}, farm.Connections, [][]string{})
}



