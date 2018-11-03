package cmd

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
	type testCaseFindArrayQuadruplet struct {
		arr    []int
		count  int
		result []int
	}

	var testCasesFindArrayQuadruplet = []testCaseFindArrayQuadruplet{
		testCaseFindArrayQuadruplet{
			[]int{2, 7, 4, 0, 9, 5, 1, 3},
			20,
			[]int{0, 4, 7, 9},
		},
		testCaseFindArrayQuadruplet{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			100,
			[]int{},
		},
		testCaseFindArrayQuadruplet{
			[]int{1, 2, 3, 4},
			12,
			[]int{},
		},
	}

	for _, item := range testCasesFindArrayQuadruplet {
		assert.Equal(
			t,
			item.result,
			FindArrayQuadruplet(item.arr, item.count),
			"should be result",
		)
	}
}

func BenchmarkFindArrayQuadruplet(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)

	for i := 0; i < b.N; i++ {
		FindArrayQuadruplet([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 100)
		FindArrayQuadruplet([]int{2, 7, 4, 0, 9, 5, 1, 3}, 20)
	}
}
