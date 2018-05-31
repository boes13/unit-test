package simple

import (
	"log"
	"testing"
)

func TestPowerOf(t *testing.T) {
	type testCase struct {
		basis    float64
		exponent int
		expected float64
	}

	testCases := make(map[string]testCase)

	testCases["case positive^zero"] = testCase{2, 0, 1}

	testCases["case positive^positive"] = testCase{2, 3, 8}

	testCases["case positive^negative"] = testCase{2, -2, 0.25}

	testCases["case negative^zero"] = testCase{-2, 0, 1}

	testCases["case negative^positive"] = testCase{-2, 3, -8}

	testCases["case negative^positive2"] = testCase{-2, 4, 16}

	testCases["case negative^negative"] = testCase{-2, -2, 0.25}

	testCases["case negative^negative2"] = testCase{-2, -1, -0.5}

	//testCases["case zero^positive"] = testCase{0, 2, 0}
	//testCases["case zero^zero"] = testCase{0, 0, ??}
	//testCases["case zero^negative"] = testCase{0, -2, error}

	for k, v := range testCases {
		log.Println("running test case:", k)
		res := PowerOf(v.basis, v.exponent)
		if res != v.expected {
			t.Errorf("expected result of %f^%d=%f, got %f", v.basis, v.exponent, v.expected, res)
		}
	}
}
