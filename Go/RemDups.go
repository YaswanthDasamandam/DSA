package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}
type Linkedlist struct {
	Head *Node
}

// Methods
// Print
// append
// Pop

func (l *Linkedlist) Print() {
	// This method Prints all the elements in the list
	current := l.Head
	for current.Next != nil {
		fmt.Printf("%v ", current.Val)
		current = current.Next
	}
	fmt.Printf("%v\n", current.Val)
}

func (l *Linkedlist) append(data int) *Linkedlist {
	// By defination it should append the node to the end node
	new_node := Node{Val: data}

	if l.Head == nil {
		l.Head = &new_node
		return l
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &new_node
		return l
	}

}

// Pop is not needed
func (l *Linkedlist) Pop() int {

	current := l.Head
	if current == nil {
		fmt.Printf("There are no element in the list")
	}
	var prev *Node = nil
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	fmt.Printf("Last previous element is %v\n", prev.Val)
	fmt.Printf("Last element is : %v\n", current.Val)
	prev.Next = nil
	value := current.Val
	current.Next = nil
	return value

}

// using Hash Maps
func (l *Linkedlist) Rem_Dups() {

	m := make(map[int]bool)

	current := l.Head
	var prev *Node = current

	m[current.Val] = true

	for current.Next != nil {
		if m[current.Val] {
			prev.Next = current.Next
		} else {
			m[current.Val] = true
		}
		prev = current
		current = current.Next
	}

	return

}

// :Todo: using runner write the Rem_Dups
func (l *Linkedlist) Rem_Dups_runner() {
	//
}

func main() {

	l := Linkedlist{}
	var pow = []int{1, 1, 2, 2, 4, 8, 16, 32, 64, 16, 128, 128}
	for _, v := range pow {
		l.append(v)
	}
	l.append(5)
	l.Print()
	l.Rem_Dups()
	l.Print()
}
