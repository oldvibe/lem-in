package ants

func Contains(Path []string, NextRoom string) bool {
	for i := 0; i < len(Path); i++ {
		if Path[i] == NextRoom {
			return true
		}
	}
	return false
}