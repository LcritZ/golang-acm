package basic

func RectCover( number int ) int {
	// write code here
	if number < 1 {
		return 0
	}
	var res int = 1
	if number == 1 {
		return res
	}
	if number == 2 {
		return res+1
	}
	if number >= 3 {
		res =  RectCover(number-1) + RectCover(number-2)
	}
	return res
}
