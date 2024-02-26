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

func main() {

	x := Node{Val: 5}
	y := Node{Val: 6}
	z := Node{Val: 7}

	l := Linkedlist{Head: &x}

	x.Next = &y
	y.Next = &z

	fmt.Printf("%v\n", l.Head)

	l.Print()
	l.append(25)
	l.Print()
	l.Pop()
	l.Print()
}
