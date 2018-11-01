package main

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
