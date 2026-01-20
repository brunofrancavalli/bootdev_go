package main

func getLast[T any](s []T) T {
	var returnValue T

	if len(s) > 0 {
		returnValue = s[len(s)-1]
	}

	return returnValue
}
