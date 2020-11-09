package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var (
	// ErrNilList returns if list is empty.
	ErrNilList = errors.New("List is nil")
	// ErrIndexList returns if index of list is out of bounds
	ErrIndexList = errors.New("Index of list is out of bounds")
)

var (
	menuText = `
=========>MENU<==========
1 - print all list.
2 - print item
3 - add item
4 - remove item
c - clear console
q - exit
=========================
`
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
	list := InitList()
	menu(list)
}

// menu ...
func menu(list *List) {
	var input string

	for input != "q" {

		fmt.Println(menuText)
		fmt.Print(">>> ")
		fmt.Scan(&input)

		switch input {

		case "1":
			list.PrintList()
		case "2":
			list.PrintListByIndex(readIndex())
		case "3":
			list.AddItem(readItem())
		case "4":
			list.RemoveItemByIndex(readIndex())
		case "c":
			clearConsoleWindows()
		case "q":
			fmt.Println("Exit...")
		default:
			fmt.Println("Wrong command")
		}

	}

}

func clearConsoleWindows() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readIndex() int {
	var in int
	fmt.Print("Enter index (from 0 to lenght-1): ")
	fmt.Scan(&in)

	return in
}

func readItem() int {
	var in int
	fmt.Print("Enter number (default - 0): ")
	fmt.Scan(&in)

	return in
}
