package acm

/**

 */
func Fibonacci( n int ) int {
	// write code here
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	f1 := 1
	f2 := 1
	f3 := 2
	for n >= 3 {
		f3 = f2 + f1
		f1 = f2
		f2 = f3
		n--
	}
	return f3
}
