package main

import (
	"errors"
	"fmt"
)

var (
	// ErrNilList returns if list is empty.
	ErrNilList = errors.New("List is nil")
	// ErrIndexList returns if index of list is out of bounds
	ErrIndexList = errors.New("Index of list is out of bounds")
)

// ListNode structure of the list node.
type ListNode struct {
	value    int       // value of node
	next     *ListNode // next node
	previous *ListNode // previous node
}

// List structure of the list.
type List struct {
	len  int       // lenght of the list
	head *ListNode // head of list
	tail *ListNode // tail of list
}

// InitList returns new list.
func InitList() *List {
	return &List{}
}

// AddItem adds item to list from the end.
func (l *List) AddItem(value int) error {
	li := &ListNode{value, nil, nil}
	if l.head == nil {
		l.head = li
	} else {
		currentNode := l.tail
		currentNode.next = li
		li.previous = l.tail
	}
	l.tail = li
	l.len++

	return nil
}

// PrintList prints all items of list.
func (l *List) PrintList() error {
	currentNode := l.head
	if currentNode == nil {
		fmt.Println("List is empty.")
		return ErrNilList
	}

	list := l.head
	fmt.Print("List: ")
	for list != nil {
		fmt.Printf("%+v -> ", list.value)
		list = list.next
	}
	fmt.Println()
	return nil
}

// Len returns lenght of list.
func (l *List) Len() int {
	return l.len
}

// PrintListByIndex prints item by index (from 0 to list lenght-1).
func (l *List) PrintListByIndex(index int) error {
	if l.len == 0 {
		fmt.Println("List is empty.")
		return ErrNilList
	}
	if index < 0 || index > l.len-1 {
		fmt.Println("Index", index, "of list is out of bounds.")
		return ErrIndexList
	}

	list := l.head
	for i := 0; i < index; i++ {
		list = list.next
	}
	fmt.Printf("Item with index %d: %+v\n", index, list.value)
	return nil
}

// RemoveItemByIndex removes item of list with the specified index (from 0 to list lenght-1).
func (l *List) RemoveItemByIndex(index int) error {
	// if list is empty
	if l.len == 0 {
		fmt.Println("List is empty.")
		return ErrNilList
	}
	// if index id out of bounds
	if index < 0 || index > l.len-1 {
		fmt.Println("Index", index, "of list is out of bounds.")
		return ErrIndexList
	}
	// list by index
	list := l.head
	for i := 0; i < index; i++ {
		list = list.next
	}
	// if is first item
	if list.previous == nil {
		list.next.previous = nil
		l.head = list.next
		return nil
	}
	// if is last item
	if list.next == nil {
		list.previous.next = nil
		l.tail = list.previous
		return nil
	}
	// default case
	list.next.previous = list.previous
	list.previous.next = list.next
	list = nil

	return nil
}

func main() {
	a := InitList()
	a.AddItem(8)
	a.AddItem(6)
	a.AddItem(3)
	a.AddItem(0)
	a.AddItem(1)
	a.PrintList()
	a.RemoveItemByIndex(4)
	a.AddItem(99)
	a.AddItem(7)
	a.RemoveItemByIndex(2)
	a.PrintList()

}

// menu ...
func menu() {

}
