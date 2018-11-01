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
