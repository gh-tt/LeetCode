package main

func countSquares(matrix [][]int) int {
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				continue
			} else if i == 0 || j == 0 {
				res++
			} else {
				matrix[i][j] = min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1])) + 1
				res += matrix[i][j]
			}
		}
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
