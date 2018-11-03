package cmd

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	var result bool
	var invalidValues = [][2]string{
		[2]string{"aa", "ab"},
		[2]string{"ba", "ab"},
		[2]string{"ad", "ac"},
		[2]string{"ae", "ad"},
		[2]string{"aa", "a*b"},
	}

	for _, value := range invalidValues {
		result = IsMatch(value[0], value[1])

		if result {
			t.Error("Should be invalid ", value)
		}
	}

	var validValues = [][2]string{
		[2]string{"aa", "aa"},
		[2]string{"", ""},
		[2]string{"ad", "ac*d"},
		[2]string{"ad", "a.*d"},
		[2]string{"ae", "a*d*e*"},
		[2]string{"aa", "a*b*a"},
		[2]string{"aa", ".."},
		[2]string{"aa", "..*"},
		[2]string{"", ".*"},
		[2]string{"abcdefgz", ".*"},
		[2]string{"abcdefgz", ".*e*"},
		[2]string{"*", "..*"},
		[2]string{".", "..*"},
	}

	for _, value := range validValues {
		result = IsMatch(value[0], value[1])

		if !result {
			t.Error("Should be valid ", value)
		}
	}
}
