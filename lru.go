package main

//LRU缓存
import "fmt"

type LRU struct {
	capacity   int
	head, tail *Node
	remember   map[int]*Node
}
type Node struct {
	key, value int
	prev, next *Node
}

func CreateLRU(capacity int) LRU {
	return LRU{
		capacity: capacity,
		head:     nil,
		tail:     nil,
		remember: make(map[int]*Node),
	}
}

func (L *LRU) Get(key int) int {
	if node, ok := L.remember[key]; !ok {
		return -1
	} else if node == L.tail {
		return node.value
	} else {
		if node.prev != nil {
			node.prev.next = node.next
		} else {
			L.head = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
		node.prev = L.tail
		node.next = nil
		if L.tail != nil {
			L.tail.next = node
		}
		L.tail = node
		if L.head == nil {
			L.head = node
		}
		return node.value
	}
}

func (L *LRU) Put(key int, value int) {
	if L.head == nil {
		L.head = &Node{
			key:   key,
			value: value,
			prev:  nil,
			next:  nil,
		}
		L.remember[key] = L.head
		L.tail = L.head
		return
	}
	if _, ok := L.remember[key]; !ok {
		if len(L.remember) == L.capacity {
			delete(L.remember, L.head.key)
			L.head = L.head.next
			L.head.prev = nil
			L.tail.next = &Node{
				key:   key,
				value: value,
				prev:  L.tail,
				next:  nil,
			}
			L.tail = L.tail.next
			L.remember[key] = L.tail
		} else {
			L.tail.next = &Node{
				key:   key,
				value: value,
				prev:  L.tail,
				next:  nil,
			}
			L.tail = L.tail.next
			L.remember[key] = L.tail
		}
	} else {
		return
	}
}

func (L *LRU) GetThenPrint(key int) {
	L.Get(key)
	for i := L.tail; i != nil; i = i.prev {
		fmt.Printf("%d\t", i.value)
	}
	fmt.Println()
}
func (L *LRU) PutThenPrint(key int, value int) {
	L.Put(key, value)
	for i := L.tail; i != nil; i = i.prev {
		fmt.Printf("%d\t", i.value)
	}
	fmt.Println()
}
func (L *LRU) Print() {
	for i := L.tail; i != nil; i = i.prev {
		fmt.Printf("%d\t", i.value)
	}
	fmt.Println()
}
func main() {
	L := CreateLRU(3)
	L.Put(1, 1) //键值为1，数值为3
	L.Put(2, 2)
	L.Put(3, 3)
	L.Print()
	L.PutThenPrint(4, 4)
	L.GetThenPrint(2)
	L.GetThenPrint(4)
	for i := 0; i < 10; i++ {
		L.PutThenPrint(i, i)
	}
}
