package simple

import "testing"

type testCase struct {
	basis    float64
	exponent int
	expected float64
}

func TestPowerOf(t *testing.T) {
	testCases := []testCase{
		//positive basis, positive or zero exponent
		{2, 0, 1},
		{2, 3, 8},
		{2, 4, 16},

		//negative basis, negative or zero exponent
		{-2, 0, 1},
		{-2, 3, -8},
		{-2, 4, 16},

		////positive basis, negative exponent
		//{2, -1, 0.5},
		//{2, -2, 0.25},
		//
		////negative basis, negative exponent
		//{-2, -1, -0.5},
		//{-2, -2, 0.25},

		////zero basis
		//{0, 2, 0},
		//{0, 0, 1},
		//{0, -2, 0},
	}

	for _, v := range testCases {
		res := PowerOf(v.basis, v.exponent)
		if res != v.expected {
			t.Errorf("expected result of %f^%d=%f, got %f", v.basis, v.exponent, v.expected, res)
		}
	}
}
