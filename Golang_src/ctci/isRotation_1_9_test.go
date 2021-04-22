package ctci

import "testing"

func TestIsRotation(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"waterbottl", "erbottlewat", false},
		{"", "", true},
		{" ", " ", true},
		{"waterbottle", "waterbottle", true},
		{"affdfdsfsff", "waterbottle", false},
	}

	for _, c := range cases {
		actual := IsRotation(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs: %s, %s, Expected: %v, actual: %v\n", c.input1, c.input2, c.expected, actual)
		}
	}
}
