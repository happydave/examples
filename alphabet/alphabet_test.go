package alphabet

import (
	"testing"
)

// ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ

func TestCountTicks(t *testing.T) {
	testCountTicksHelper(t, "AZ", 1)
	testCountTicksHelper(t, "AZB", 3)
	testCountTicksHelper(t, "AYB", 5)
	testCountTicksHelper(t, "ABYC", 8)
	testCountTicksHelper(t, "BXCDC", 12)
	testCountTicksHelper(t, "bXvDqd", 41)
	testCountTicksHelper(t, "AAAAAAAAAAAAAAAA", 0)
	testCountTicksHelper(t, "ZZZZZZZZZZZZZZZZ", 1)
	testCountTicksHelper(t, "MMMMMMMMMMMMMMMM", 12)
	testCountTicksHelper(t, "NNNNNNNNNNNNNNNN", 13)
}

func testCountTicksHelper(t *testing.T, input string, expectedResult int) {
	result := countTicks(input)
	if result != expectedResult {
		t.Errorf("FAILED: expected %d returned %d string %s", expectedResult, result, input)
	} else {
		t.Logf("PASSED: (%d) %s", expectedResult, input)
	}
}
