package ctci

import "testing"

func TestBasicCompress(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"aaaabbbcdddd", "a4b3c1d4"},
		{"abcccccddde", "a1b1c5d3e1"},
		{"", ""},
		{"a", "a"},
		{"abcd", "abcd"}, // compressed is longer, expect input
	}
	for _, c := range cases {
		actual := BasicCompress(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %s, actual: %s\n", c.input, c.expected, actual)
		}
	}
}
