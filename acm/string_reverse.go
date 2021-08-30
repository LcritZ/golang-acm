package acm

func StringReverse(input string ) string {
	if len(input) == 0 {
		return input
	}

	length := len(input)
	res := make([]byte, length, length)

	i := length -1
	for _, item := range input {
		ch := item
		res[i] = byte(ch)
		i --
	}
	return string(res)
}



