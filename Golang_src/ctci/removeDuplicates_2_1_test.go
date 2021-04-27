package ctci

import (
	"reflect"
	"testing"
)

// Use hashset:
// Time Complexity: O(N)
// Space Complexity:O(N)
func TestRemoveDuplicate(t *testing.T) {
	cases := []struct {
		input    *ListNode
		expected *ListNode
	}{} // FIXME: need add cases
	for _, c := range cases {
		removeDuplicate(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, c.input)
		}
	}
}
