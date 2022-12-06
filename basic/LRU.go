package basic

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	Cap int
	Keys map[int]*Node
	Head *Node
	Tail *Node
}

func ConstructLRUCache(capacity int) LRUCache {
	return LRUCache{
		Cap: capacity,
		Keys: map[int]*Node{},
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.Keys[key]; ok {
		cache.Remove(node)
		cache.Add(node)
		return node.Val
	}

	return -1
}

func (cache *LRUCache) Put(key, value int) {
	if node, ok := cache.Keys[key]; ok {
		cache.Remove(node)
		cache.Add(node)
		return
	}

	if len(cache.Keys) == cache.Cap {
		delete(cache.Keys, cache.Tail.Key)
		cache.Remove(cache.Tail)
	}
	node := &Node{
		Key: key,
		Val: value,
	}
	cache.Keys[key] = node
	cache.Add(node)

	return
}

func (cache *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = cache.Head
	if cache.Head != nil {
		cache.Head.Prev = node
	}
	cache.Head = node
	if cache.Tail == nil {
		cache.Tail = node
		cache.Tail.Next = nil
	}
}

func (cache *LRUCache) Remove(node *Node)  {
	if node == cache.Head {
		cache.Head = node.Next
		if node.Next != nil {
			node.Next.Prev = nil
		}
		node.Next = nil

		return
	}
	if node == cache.Tail {
		cache.Tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
