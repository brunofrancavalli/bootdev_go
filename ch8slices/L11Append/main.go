package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	filteredCost := []float64{}
	for index := 0; index < len(costs); index++ {
		if costs[index].day == day {
			filteredCost = append(filteredCost, costs[index].value)
		}
	}

	return filteredCost
}