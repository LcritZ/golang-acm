package leetcode

import (
	"math"
	"strconv"
)

func ReverseInt(x int) int {
	flag := false
	if x < 0  {
		flag = true
		x = -x
	}
	strNum := strconv.Itoa(x)
	base := 1
	ansNum := 0
	for i := 0; i<len(strNum); i++ {
		if i == 0 && strNum[i] == '0' {
			continue
		}
		if ansNum < math.MinInt32/10 || ansNum > math.MaxInt32/10 {
			return 0
		}
		ansNum +=  int((strNum[i])-48) *base
		base *= 10
	}
	if ansNum >  math.MaxInt32 {
		return 0
	}
	if flag {
		ansNum = -ansNum
	}
	return ansNum
}

func GF_reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

func reverseInt2(x int) int {
	res := 0
	for x !=0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x%10
		x /= 10
		res = 10*res+digit
	}

	return res
}
