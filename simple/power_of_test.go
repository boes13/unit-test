package simple

import "testing"

type testCase struct {
	basis int
	exponent int
	expected float64
}

func TestPowerOf(t *testing.T) {
	testCases := []testCase{
		//positive basis, positive or zero exponent
		{2, 0, 1},
		{2,3, 8},
		{2,4, 16},

		//negative basis, negative or zero exponent
		{-2, 0, 1},
		{-2, 3, -8},
		{-2, 4, 16},

		//positive basis, negative exponent
		{2, -1, 0.5},
		{2, -2, 0.25},

	}
}