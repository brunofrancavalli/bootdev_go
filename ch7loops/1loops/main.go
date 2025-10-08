package main

func bulkSend(numMessages int) float64 {
	total := 0.0
	for count := 0; count < numMessages; count++ {
		total += 1.0 + float64(count)/float64(100)
	}

	return total
}
