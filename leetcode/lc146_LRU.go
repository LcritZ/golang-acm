package leetcode

//type CacheNode struct {
//    Key int
//    Val int
//    Prev *CacheNode
//    Next *CacheNode
//}
//
//type LRUCache struct {
//    Cap int
//    Keys map[int]*CacheNode
//    Head *CacheNode
//    Tail *CacheNode
//}
//
//func ConstructorLRU(capacity int) LRUCache {
//    return LRUCache{
//        Cap: capacity,
//        Keys: make(map[int]*CacheNode, capacity),
//    }
//}
//
//func (c *LRUCache) Get(key int) int {
//    if node, ok := c.Keys[key]; ok {
//        c.Remove(node)
//        c.Add(node)
//        return node.Val
//    } else {
//        return -1
//    }
//}
//
//func (c *LRUCache) Put(key, val int)  {
//    if node, ok := c.Keys[key]; ok {
//        c.Remove(node)
//        node.Val = val
//        c.Add(node)
//    } else {
//        node2 := &CacheNode{
//            Key: key,
//            Val: val,
//        }
//        if len(c.Keys) == c.Cap {
//            fmt.Println("remove tail", c.Tail.Val)
//            c.Remove(c.Tail)
//        }
//        c.Add(node2)
//    }
//
//}
//
//func (c *LRUCache) Add(node *CacheNode) {
//    c.Keys[node.Key] = node
//    node.Next = c.Head
//    if c.Head != nil {
//        c.Head.Prev = node
//    }
//    c.Head = node
//    if c.Tail == nil {
//        c.Tail = c.Head
//    }
//    return
//}
//
//func (c *LRUCache) Remove(node *CacheNode)  {
//    if node == c.Head {
//        c.Head = node.Next
//        if c.Head != nil && c.Head.Next != nil {
//            c.Head.Next.Prev = c.Head
//        }
//        node.Next = nil
//        delete(c.Keys, node.Key)
//        return
//    }
//    if node == c.Tail {
//        fmt.Println("tail---")
//        if node.Prev != nil {
//            node.Prev.Next = nil
//            c.Tail = node.Prev
//        } else {
//            c.Tail = nil
//        }
//        node.Prev = nil
//        delete(c.Keys, node.Key)
//        return
//    }
//    node.Prev.Next = node.Next
//    node.Next.Prev = node.Prev
//    node.Prev = nil
//    node.Next = nil
//    delete(c.Keys, node.Key)
//    return
//}
//

type LRUCache struct {
    capacity int
    m map[int]*Node
    head, tail *Node
}

type Node struct {
    Key int
    Value int
    Pre, Next *Node
}


func ConstructorLRU(capacity int) LRUCache {
    head, tail := &Node{}, &Node{}
    head.Next = tail
    tail.Pre = head
    return LRUCache{
        capacity: capacity,
        m: map[int]*Node{},
        head: head,
        tail: tail,
    }
}

func (this *LRUCache) Get(key int) int {
    if v, ok := this.m[key]; ok {
        this.moveToHead(v)
        return v.Value
    }
    return -1
}

func (this *LRUCache) moveToHead(node *Node) {
    this.deleteNode(node)
    this.addToHead(node)
}

func (this *LRUCache) deleteNode(node *Node) {
    node.Pre.Next = node.Next
    node.Next.Pre = node.Pre
}

func (this *LRUCache) removeTail() int {
    node := this.tail.Pre
    this.deleteNode(node)
    return node.Key
}

func (this *LRUCache) addToHead(node *Node) {
    this.head.Next.Pre = node
    node.Next = this.head.Next
    node.Pre = this.head
    this.head.Next = node
}

func (this *LRUCache) Put(key int, value int)  {
    if v, ok := this.m[key]; ok {
        v.Value = value
        this.moveToHead(v)
        return
    }

    if this.capacity == len(this.m) {
        rmKey := this.removeTail()
        delete(this.m, rmKey)
    }

    newNode := &Node{Key: key, Value: value}
    this.addToHead(newNode)
    this.m[key] = newNode
}
