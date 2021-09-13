package basic

func JumpFloor( number int ) int {
	// write code here
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}

	f1 := 1
	f2 := 2
	f3 := 3
	for number >= 3 {
		f3 = f2 + f1
		f1 = f2
		f2 = f3
		number--
	}
	return f3
}
