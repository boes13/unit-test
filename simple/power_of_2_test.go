package simple

import "testing"

type testCase2 struct {
	basis    float64
	exponent int
	expected float64
	error    bool
}

func TestPowerOf_3(t *testing.T) {
	testCases := []testCase2{
		//positive basis, positive or zero exponent
		{2, 0, 1, false},
		{2, 3, 8, false},
		{2, 4, 16, false},

		//negative basis, negative or zero exponent
		{-2, 0, 1, false},
		{-2, 3, -8, false},
		{-2, 4, 16, false},

		//positive basis, negative exponent
		{2, -1, 0.5, false},
		{2, -2, 0.25, false},

		//negative basis, negative exponent
		{-2, -1, -0.5, false},
		{-2, -2, 0.25, false},

		//zero basis
		{0, 2, 0, false},
		{0, 0, 1, true},
		{0, -2, 0, true},
	}

	for _, v := range testCases {
		res, err := PowerOf_3(v.basis, v.exponent)
		if v.error && err == nil {
			t.Fatalf("expected error on calculation %f^%d, got no error", v.basis, v.exponent)
		}
		if !v.error && res != v.expected {
			t.Errorf("expected result of %f^%d=%f, got %f", v.basis, v.exponent, v.expected, res)
		}
	}
}