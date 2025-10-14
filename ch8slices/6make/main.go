package main

func getMessageCosts(messages []string) []float64 {
	costSlice := make([]float64, len(messages))

	for index := 0; index < len(messages); index++ {
		costSlice[index] = float64(len(messages[index])) * float64(0.01)
	}

	return costSlice
}
