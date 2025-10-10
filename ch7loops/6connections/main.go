package main

func countConnections(groupSize int) int {
	value := 0

	if groupSize > 0 {
		if groupSize-1 > 0 {
			value = countConnections(groupSize - 1)
		}

		value += (groupSize - 1)
	}

	return value

}
