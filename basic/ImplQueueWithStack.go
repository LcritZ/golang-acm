package basic

var stack1 [] int
var stack2 [] int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int{
	var res = stack1[0]
	stack1 = stack1[1:]
	return res
}
