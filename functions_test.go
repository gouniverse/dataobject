package dataobject

import (
	"testing"
)

func TestToString(t *testing.T) {
	inputs := [][]any{
		{1, "1"},
		{"2", "2"},
		{true, "true"},
		{false, "false"},
		{0.123, "0.1230"}, // precission 4
		{nil, ""},
	}

	for _, v := range inputs {
		result := toString(v[0])
		expected := v[1]

		if result != expected {
			t.Error("Result", result, "DOES NOT match expected:", expected)
		}
	}
}
