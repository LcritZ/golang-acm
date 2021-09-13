package basic

import "fmt"

var Post []int
var Res []byte

/**
	     1
	 2     5
   3  4	 6   7
8
var pre = []int{1,2,3,8,4,5,6,7}
var vin = []int{8,3,2,4,1,6,5,7}
index

pre[1:index+1]
vin[0:index]
2,3,8,4
8,3,2,4

pre[index+1:]
vin[index+1:]

234
324
index=1
pre[1:2]  3
vin[0:1]  3

pre[2:] 4
vin[2:] 4

 */

func GetPostOrder(pre []int, in []int) {
	fmt.Println(pre)
	fmt.Println(in)
	fmt.Println(Post, "--")
	if len(pre) == 0 {
		return
	}
	root := pre[0]
	if len(pre) == 1 {
		Post = append(Post, root)
		return
	}
	var index int
	for i := 0; i < len(pre); i++ {
		if in[i] == root {
			index = i
			break
		}
	}
	fmt.Println(index, "----")
	if len(pre) > 1 && index+1 < len(pre) && index < len(in) {
		GetPostOrder(pre[1:index+1], in[:index])
	}
	if index+1 < len(pre) && index+1 < len(in) {
		GetPostOrder(pre[index+1:], in[index+1:])
	}
	Post = append(Post, root)
}
//
var pre = []int{1,2,3,8,4,5,6,7}
var vin = []int{8,3,2,4,1,6,5,7}
//var pre = []byte{'A','B', 'D','E', 'G', 'C', 'F'}
//var vin = []byte{'D','B','G','E','A','C','F'}

//func PostOrder(pre []int, vin []int)  {
//
//}

func GetPost(root int, start, end int) {
	//fmt.Println(root, start, end, "--")
	if(start > end) {
		return
	}
	i := start
	for i < end && vin[i] != pre[root] {
		i++
	}
	fmt.Println(pre[start:end+1])
	GetPost(root + 1,           start,           i - 1);
	GetPost(root + 1 + i-start, i + 1, end)
	Post = append(Post, pre[root])
}
