package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralFromCenter(t *testing.T) {
	assert.Equal(
		t,
		SpiralFromCenter(
			[][]int{
				[]int{9, 2, 3},
				[]int{8, 1, 4},
				[]int{7, 6, 5},
			},
		),
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		"they should be equal",
	)
}

func BenchmarkSpiralFromCenter(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)

	for i := 0; i < b.N; i++ {
		SpiralFromCenter(
			[][]int{
				[]int{9, 2, 3},
				[]int{8, 1, 4},
				[]int{7, 6, 5},
			},
		)
	}
}
