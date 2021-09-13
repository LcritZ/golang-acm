package basic

func NumberOf1( n int ) int {
	// write code here
	//if n == 0 {
	//	return 0
	//}
	//var resOne int = 0
	//var resZero int = 0
	//var temp int = 0
	//if n > 0 {
	//	temp = n
	//} else if n < 0 {
	//	temp = -n
	//}
	//for temp > 0 {
	//	tail := temp%2
	//	if tail == 1 {
	//		resOne++
	//	} else {
	//		resZero++
	//	}
	//	temp = temp/2
	//}
	//var res int
	//var isodd bool
	//if n > 0 {
	//	res = resOne
	//} else if n < 0 {
	//	if n%2 == 0 {
	//		isodd = true
	//	}
	//	res = 32 - resOne
	//	if !isodd {
	//		res += 1
	//	}
	//}

	//第二种
	if n == 0 {
		return 0
	}
	var res int = 0
	var temp int = n
	if n < 0 {
		temp = -n - 1
	}
	for temp != 0 {
		temp = temp&(temp-1)
		res ++
	}
	if n < 0 {
		res = 32-res
	}

	return res
}