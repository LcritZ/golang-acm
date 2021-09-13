package basic

import "sync"

type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)

	// 栈中元素数量+1
	stack.size = stack.size + 1
}

// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}