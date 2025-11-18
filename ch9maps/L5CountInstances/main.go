package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, messagedUser := range messagedUsers {
		if _, ok := validUsers[messagedUser]; ok {
			validUsers[messagedUser] = validUsers[messagedUser] + 1
		}
	}
}
