package basic

func ReOrderArray( array []int ) []int {
	// write code here
	var length int = len(array)
	if length== 0 {
		return array
	}
	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if (array[j]%2 == 0) && (array[j+1]%2 == 1) {
				var temp int = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}

	return array
}
