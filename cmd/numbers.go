package cmd

import "math"

// GetGreatestCommonDivisor find by division
func GetGreatestCommonDivisor(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	return a + b
}

// GetGreatestCommonDivisor2 find by decrease
func GetGreatestCommonDivisor2(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

// IsPrimeNumber desc
func IsPrimeNumber(n int) bool {
	if n < 4 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	end := int(math.Sqrt(float64(n)))

	for i := 5; i <= end; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
