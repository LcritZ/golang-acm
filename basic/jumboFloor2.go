package basic

func JumpFloorII( number int ) int {
	// write code here
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	var res int = 1
	for number >= 2 {
		res = 2*res
		number --
	}
	return res
}
