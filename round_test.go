package asset

import "testing"

func TestRoundUp(t *testing.T) {
	for _, n := range roundUpTests {
		if RoundUp(n.input, n.digits) != n.output {
			t.Error("Failed to round up")
		}
	}
}

func TestRoundDown(t *testing.T) {
	for _, n := range roundDownTests {
		if RoundDown(n.input, n.digits) != n.output {
			t.Error("Failed to round down")
		}
	}
}
