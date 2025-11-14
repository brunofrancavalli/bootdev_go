package main

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}

	for row := 0; row < rows; row++ {
		rowTemp := []int{}
		for col := 0; col < cols; col++ {
			rowTemp = append(rowTemp, col*row)
		}
		matrix = append(matrix, rowTemp)
	}

	return matrix
}
