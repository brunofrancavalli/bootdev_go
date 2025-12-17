package main

func sum(nums ...int) int {
	value := 0

	for index := 0; index < len(nums); index++ {
		value += nums[index]
	}

	return value
}
