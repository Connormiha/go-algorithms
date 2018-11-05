package cmd

import (
	"testing"
)

var testCasesCommonDivisor = [][3]int{
	{10, 50, 10},
	{12, 6, 6},
	{100000000, 5000, 5000},
	{124, 10, 2},
	{124, 100, 4},
	{555, 30, 15},
	{45, 15000, 15},
}

func TestGetGreatestCommonDivisor(t *testing.T) {
	for _, value := range testCasesCommonDivisor {
		result := GetGreatestCommonDivisor(value[0], value[1])
		if result != value[2] {
			t.Error("Should be valid ", result, " != ", value[2])
		}
	}
}
