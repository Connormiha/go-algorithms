package cmd

import "sort"

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// MaxSumSubArray desc
func MaxSumSubArray(arr []int) int {
	totalMax := arr[0]
	currentMax := totalMax

	for i := 1; i < len(arr); i++ {
		currentMax = maxInt(arr[i], currentMax+arr[i])
		totalMax = maxInt(currentMax, totalMax)
	}

	return totalMax
}

// FindArrayQuadruplet desc
func FindArrayQuadruplet(arr []int, sum int) []int {
	length := len(arr)

	if length < 4 {
		return []int{}
	}

	sort.Ints(arr)

	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			currentSum := sum - arr[i] - arr[j]
			start := j + 1
			end := length - 1

			for start < end {
				startNum := arr[start]
				endNum := arr[end]

				totalSum := currentSum - startNum - endNum

				if totalSum == 0 {
					return []int{arr[i], arr[j], startNum, endNum}
				}

				if totalSum > 0 {
					start++
				} else {
					end--
				}
			}
		}
	}

	return []int{}
}
