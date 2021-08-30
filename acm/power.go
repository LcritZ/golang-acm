package acm

func Power( base float64 ,  exponent int ) float64 {
	// write code here
	if int(base) == 0 {
		return 0
	}

	if exponent == 0 {
		return 1
	}

	var temp int = exponent
	if exponent < 0 {
		temp = -exponent
	}

	var res float64 = 1.0
	for temp >= 1 {
		res = res*base
		temp--
	}
	if exponent < 0 {
		res = 1/res
	}
	return res

}