package basic


//func compare(str1, str2 string) bool {
//	if len(str1) == 0 || len(str1) != len(str2) {
//		return false
//	}
//	arr1 := []byte(str1)
//	arr2 := []byte(str2)
//	index := 1
//	flag := false
//	for i := 0; i < len(arr1); i++ {
//		j := 0
//		if arr1[i] == arr2[j] {
//			for ; i < len(arr1) && j < len(arr2); i++ j++ {
//				if arr1[i] != arr2[j] {
//					break
//				}
//				if i == len(arr1) -1 {
//					flag = true
//					index = j
//				}
//				continue
//			}
//		}
//		if flag {
//			break
//		}
//		continue
//	}
//	j := index-1
//	if flag {
//		for i := 0; i < len(arr1)-index && j < len(arr1); i++, j++ {
//			if arr1[i] != arr2[j] {
//				return false
//			}
//			continue
//		}
//	}
//	if j == len(arr1) -1 {
//		return true
//	}
//	return false
//}

//
//func main() {
//	//a := 0
//	//fmt.Scan(&a)
//	//fmt.Printf("%d\n", a)
//	fmt.Printf("Hello World!\n");
//}
