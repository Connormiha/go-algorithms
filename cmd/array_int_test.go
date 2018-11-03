package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubArray(t *testing.T) {
	assert.Equal(
		t,
		MaxSumSubArray([]int{10, -11, 12}),
		12,
		"should be 12",
	)
}

func TestFindArrayQuadruplet(t *testing.T) {
	type testCase struct {
		arr    []int
		count  int
		result []int
	}

	testCases := []testCase{
		testCase{
			[]int{2, 7, 4, 0, 9, 5, 1, 3},
			20,
			[]int{0, 4, 7, 9},
		},
		testCase{
			[]int{1, 2, 3, 4},
			12,
			[]int{},
		},
	}

	for _, item := range testCases {
		assert.Equal(
			t,
			item.result,
			FindArrayQuadruplet(item.arr, item.count),
			"should be result",
		)
	}
}
