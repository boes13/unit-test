package simple

import "fmt"

// PowerOf returns the power of basis ~ basis^exponent.
func PowerOf(basis float64, exponent int) float64 {
	retval := float64(1)

	for i := 0; i < exponent; i++ {
		retval *= basis
	}

	return retval
}

func PowerOf_2(basis float64, exponent int) float64 {
	retval := float64(1)

	negativeExponent := false
	if exponent < 0 {
		negativeExponent = true
		exponent *= -1
	}

	for i := 0; i < exponent; i++ {
		if negativeExponent {
			retval *= (1 / basis)
		} else {
			retval *= basis
		}
	}

	return retval
}

func PowerOf_3(basis float64, exponent int) (float64, error) {
	if basis == 0 && exponent <= 0 {
		return 0, fmt.Errorf("uncalculatable operation")
	}

	retval := float64(1)
	negativeExponent := false
	if exponent < 0 {
		negativeExponent = true
		exponent *= -1
	}

	for i := 0; i < exponent; i++ {
		if negativeExponent {
			retval *= (1 / basis)
		} else {
			retval *= basis
		}
	}

	return retval, nil
}
