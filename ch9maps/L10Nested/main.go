package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCount := make(map[rune]map[string]int)
	for _, name := range names {
		firstRune := []rune(name)[0]
		nameRune, ok := nameCount[firstRune]
		if !ok {
			nameRune = make(map[string]int)
			nameCount[firstRune] = nameRune
		}
		nameRune[name]++
	}
	return nameCount
}
