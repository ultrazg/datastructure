package main

import "fmt"

type node struct {
	name  string
	no    int
	score int
	next  *node
}

func (n *node) newNode(name string, no, score int) *node {
	return &node{
		name:  name,
		no:    no,
		score: score,
		next:  nil,
	}
}

type linkedList struct {
	first *node
	last  *node
}

func (linkedList *linkedList) isEmpty() bool {
	return linkedList.first == nil
}

func (linkedList *linkedList) print() {
	current := linkedList.first

	// 如果当前节点存在
	for current != nil {
		fmt.Printf("Name: %s, No: %d, Score: %d\n", current.name, current.no, current.score)
		current = current.next
	}
}

func (linkedList *linkedList) insert(name string, no, score int) {
	newNode := &node{
		name:  name,
		no:    no,
		score: score,
		next:  nil,
	}

	// 如果链表为空
	if linkedList.isEmpty() {
		linkedList.first = newNode
		linkedList.last = newNode
	} else {
		linkedList.last.next = newNode
		linkedList.last = newNode
	}
}

func (linkedList *linkedList) delFirst() {
	if linkedList.isEmpty() {
		println("当前链表为空")
		return
	}

	linkedList.first = linkedList.first.next
}

func (linkedList *linkedList) delLast() {
	if linkedList.isEmpty() {
		println("当前链表为空")
		return
	}

	newNode := linkedList.first
	// 遍历链表，直到找到倒数第二个节点
	for newNode.next != linkedList.last {
		newNode = newNode.next
	}

	// 将倒数第二个节点的 next 指针更新为 nil
	newNode.next = linkedList.last.next
	linkedList.last = newNode
}

func (linkedList *linkedList) delete(index int) {
	if linkedList.isEmpty() {
		println("当前链表为空")
		return
	}

	// 删除第一个节点
	if index == 0 {
		linkedList.first = linkedList.first.next

		if linkedList.first.next == nil {
			linkedList.last = nil
		}

		return
	}

	// 删除中间节点或者末尾节点
	current := linkedList.first
	for i := 0; i < index-1; i++ {
		current = current.next

		if current == nil || current.next == nil {
			println("索引超出链表的范围")

			return
		}
	}

	// 将前一个节点的 next 指针指向要删除节点的下一个节点
	current.next = current.next.next

	// 如果删除的是末尾节点，更新 next 指针
	if current.next == nil {
		linkedList.last = current
	}
}

func main() {
	list := linkedList{}

	list.insert("张三", 1, 80)
	list.insert("李四", 2, 90)
	list.insert("王五", 3, 100)

	list.print()
	println("---")
	list.delete(2)
	list.print()
}

// linkedList:
//| first | last  |
//|-------|-------|
//| node1 | node3 |
//
//node1:
//| name | no | score | next  |
//|------|----|-------|-------|
//| 张三 | 1  | 80    | node2 |
//
//node2:
//| name | no | score | next  |
//|------|----|-------|-------|
//| 李四 | 2  | 90    | node3 |
//
//node3:
//| name | no | score | next |
//|------|----|-------|------|
//| 王五 | 3  | 100    | nil  |
