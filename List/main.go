package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	menuText = "\n=========>MENU<==========\n1 - print all list.\n2 - print item\n3 - add item\n4 - remove item\nc - clear console\nq - exit\n=========================\n"
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
func (l *List) AddItem(value int) {
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
}

// PrintList prints all items of list.
func (l *List) PrintList() {
	currentNode := l.head
	if currentNode == nil {
		fmt.Println("List is empty.")
		return
	}

	list := l.head
	fmt.Print("List: ")
	for list != nil {
		fmt.Printf("%+v -> ", list.value)
		list = list.next
	}
	fmt.Println()
}

// PrintListByIndex prints item by index (from 0 to list lenght-1).
func (l *List) PrintListByIndex(index int) {
	if l.len == 0 {
		fmt.Println("List is empty.")
		return
	}
	if index < 0 || index > l.len-1 {
		fmt.Println("Index", index, "of list is out of bounds.")
		return
	}

	list := l.head
	for i := 0; i < index; i++ {
		list = list.next
	}
	fmt.Printf("Item with index %d: %+v\n", index, list.value)
}

// RemoveItemByIndex removes item of list with the specified index (from 0 to list lenght-1).
func (l *List) RemoveItemByIndex(index int) {
	// if list is empty
	if l.len == 0 {
		fmt.Println("List is empty.")
		return
	}
	// if index id out of bounds
	if index < 0 || index > l.len-1 {
		fmt.Println("Index", index, "of list is out of bounds.")
		return
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
		l.len--
		return
	} else if list.next == nil { // if is last item
		list.previous.next = nil
		l.tail = list.previous
		l.len--
		return
	} else { // default case
		list.next.previous = list.previous
		list.previous.next = list.next
		l.len--
	}

}

func main() {
	list := InitList() // Initialize list
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
