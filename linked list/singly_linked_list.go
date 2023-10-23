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

func (linkedList *linkedList) add(name string, no, score int) {
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

func (linkedList *linkedList) insert(index int, n *node) {
	newNode := &node{
		name:  n.name,
		no:    n.no,
		score: n.score,
		next:  nil,
	}

	if linkedList.isEmpty() {
		linkedList.first = newNode
		linkedList.last = newNode

		return
	}

	current := linkedList.first
	for i := 0; i < index-1; i++ {
		current = current.next

		if current == nil || current.next == nil {
			println("索引超出链表的范围")

			return
		}
	}

	if current.next.next != nil {
		newNode.next = current.next
		current.next = newNode
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

func concatenation(list1, list2 *linkedList) *linkedList {
	if list1.isEmpty() {
		list1.first = list2.first
		list1.last = list2.last
	} else {
		list1.last.next = list2.first
		list1.last = list2.last
	}

	list2.first = nil
	list2.last = nil

	return list1
}

func main() {
	students := linkedList{}
	foods := linkedList{}

	students.add("张三", 1, 80)
	students.add("李四", 2, 90)
	students.add("王五", 3, 100)

	foods.add("苹果", 1, 89)
	foods.add("香蕉", 2, 95)
	foods.add("雪梨", 3, 100)

	students.print()
	foods.print()

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
