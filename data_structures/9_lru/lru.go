package main

import "fmt"

// DLL Node
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRU Cache
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // dummy head
	tail     *Node // dummy tail
}

// Constructor
func NewLRUCache(cap int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: cap,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// Remove a node
func (l *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// Insert at front (MRU)
func (l *LRUCache) insertFront(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

// Get value
func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.remove(node)
		l.insertFront(node)
		return node.value
	}
	return -1
}

// Put key-value
func (l *LRUCache) Put(key, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.remove(node)
		l.insertFront(node)
		return
	}

	if len(l.cache) == l.capacity {
		// Remove LRU (before tail)
		lru := l.tail.prev
		l.remove(lru)
		delete(l.cache, lru.key)
	}

	node := &Node{key: key, value: value}
	l.cache[key] = node
	l.insertFront(node)
}

// Main
func main() {
	lru := NewLRUCache(2)

	lru.Put(1, 10)
	lru.Put(2, 20)

	fmt.Println(lru.Get(1)) // 10

	lru.Put(3, 30)          // evicts key 2
	fmt.Println(lru.Get(2)) // -1

	lru.Put(4, 40)          // evicts key 1
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 30
	fmt.Println(lru.Get(4)) // 40
}
