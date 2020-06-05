package main

import (
	"fmt"
)

func main() {
	re := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	fmt.Println(re)
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	rows, columns := len(matrix), len(matrix[0])
	index := 0
	res := make([]int, rows*columns)
	left, right, bottom, top := 0, columns-1, rows-1, 0

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}

		for j := top + 1; j <= bottom; j++ {
			res[index] = matrix[j][right]
			index++
		}

		if left < right && top < bottom {
			for k := right - 1; k >= left; k-- {
				res[index] = matrix[bottom][k]
				index++
			}

			for row := bottom - 1; row > top; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}
	return res
}
