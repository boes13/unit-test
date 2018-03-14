package simple

// PowerOf returns the power of basis ~ basis^exponent.
func PowerOf(basis float64, exponent int) float64 {
	retval := float64(1)

	for i:=0; i<exponent; i++ {
		retval *= basis
	}

	return retval
}