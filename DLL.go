// 双向链表
package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}
type DLL struct {
	head *Node
	tail *Node
}

func (p *DLL) InsertNode(data int) { //增
	newNode := &Node{
		data: data,
	}
	if p.tail == nil {
		p.head = newNode
		p.tail = newNode
	}
	newNode.prev = p.tail
	p.tail.next = newNode
	p.tail = newNode
}
func (p *DLL) DeleteNode(data int) { //删
	if p.head == nil {
		fmt.Println("DLL is empty")
		return
	}
	for n := p.tail; n.prev != nil; n = n.prev { //查
		if n.data == data {
			n.prev.next = n.next
			n.next.prev = n.prev
			return
		}
		if p.head.data == data {
			p.head = p.head.next
			p.head.prev = nil
			return
		}
	}
	fmt.Println("data not found")
}
func (p *DLL) UpdateNode(old_data int, new_data int) { //改
	if p.head == nil {
		fmt.Println("DLL is empty")
		return
	}
	for n := p.tail; n.prev != nil; n = n.prev { //查
		if n.data == old_data {
			n.data = new_data
			return
		}
		if n.data == old_data {
			n.data = new_data
			return
		}
	}
	fmt.Println("data not found")
}
func (p *DLL) ShowDLL() {
	if p.head == nil {
		fmt.Println("DLL is empty")
		return
	}
	var n *Node
	for n = p.head; n.next != nil; n = n.next {
		fmt.Printf("%d\t", n.data)
	}
	fmt.Printf("%d\n", n.data)
}
func main() {
	var p = &DLL{nil, nil}
	p.InsertNode(1)
	p.InsertNode(2)
	p.InsertNode(3)
	p.ShowDLL()
	p.DeleteNode(1)
	p.ShowDLL()
	p.UpdateNode(3, 9)
	p.ShowDLL()
}
