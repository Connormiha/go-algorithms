package main

import (
	"math"
)

// SpiralFromCenter foo
func SpiralFromCenter(matrix [][]int) []int {
	result := make([]int, len(matrix)*len(matrix))
	center := int(math.Floor(float64(len(matrix)) / 2))
	var i = center
	var j = center
	result[0] = matrix[i][i]

	if center == 0 {
		return result
	}

	var roundsNumber = 1
	var currentNumber = 0

	for {
		// To top
		for i > center-roundsNumber {
			i--
			currentNumber++
			result[currentNumber] = matrix[i][j]
			if i == 0 && j == 0 {
				return result
			}
		}

		// right
		for j < center+roundsNumber {
			j++
			currentNumber++
			result[currentNumber] = matrix[i][j]
		}

		// bottom
		for i < center+roundsNumber {
			i++
			currentNumber++
			result[currentNumber] = matrix[i][j]
		}

		// left
		for j > center-roundsNumber {
			j--
			currentNumber++
			result[currentNumber] = matrix[i][j]
		}

		roundsNumber++
	}
}
