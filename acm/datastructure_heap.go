package acm

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Element []int
}

// MinHeap构造方法
func NewMinHeap() *MinHeap {
	// 第一个元素仅用于结束insert中的 for 循环
	h := &MinHeap{Element: []int{math.MinInt64}}
	return h
}

// 插入数字,插入数字需要保证堆的性质
func (H *MinHeap) Insert(v int) {
	H.Element = append(H.Element, v)
	i := len(H.Element) - 1
	// 上浮
	for ; H.Element[i/2] > v; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}

	H.Element[i] = v
}

// 删除并返回最小值
func (H *MinHeap) DeleteMin() (int, error) {
	if len(H.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}
	minElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1] < H.Element[child] {
			child ++
		}
		// 下滤一层
		if lastElement > H.Element[child] {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return minElement, nil
}

// 堆的大小
func (H *MinHeap) Length() int {
	return len(H.Element) - 1
}

// 获取最小堆的最小值
func (H *MinHeap) Min() (int, error) {
	if len(H.Element) >= 1 {
		return H.Element[1], nil
	}
	return 0, fmt.Errorf("heap is empty")
}

// MinHeap格式化输出
func (H *MinHeap) String() string {
	return fmt.Sprintf("%v", H.Element[1:])
}

type Heap struct {
	Size  int
	Elems []int
}

func NewHeap(MaxSize int) *Heap {
	h := new(Heap)
	h.Elems = make([]int, MaxSize, MaxSize)
	return h
}
//最小堆
func (h *Heap) Push(x int) {
	h.Size++

	// i是要插入节点的下标
	i := h.Size
	for {
		if i <= 0 {
			break
		}

		// parent为父亲节点的下标
		parent := (i - 1) / 2
		// 如果父亲节点小于等于插入的值，则说明大小没有跌倒，可以退出
		if h.Elems[parent] <= x {
			break
		}

		// 互换当前父亲节点与要插入的值
		h.Elems[i] = h.Elems[parent]
		i = parent
	}

	h.Elems[i] = x
}

func (h *Heap) Pop() int {
	if h.Size == 0 {
		return 0
	}

	// 取出根节点
	ret := h.Elems[0]

	// 将最后一个节点的值提到根节点上
	h.Size--
	x := h.Elems[h.Size]

	i := 0
	for {
		// a，b为左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 没有左子树
		if a >= h.Size {
			break
		}

		// 有右子树，找两个子节点中较小的值
		if b < h.Size && h.Elems[b] < h.Elems[a] {
			a = b
		}

		// 父亲小直接退出
		if h.Elems[a] >= x {
			break
		}

		// 交换
		h.Elems[i] = h.Elems[a]
		i = a
	}

	h.Elems[i] = x
	return ret
}

func (h *Heap) Display() {
	fmt.Printf("Size:%d,Elems:%#v\n", h.Size, h.Elems[0:h.Size])
}
