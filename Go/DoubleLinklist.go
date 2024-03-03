package main

import "fmt"

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

type DLinkedlist struct {
	Head *Node
}

// Methods
// Print

func (l *DLinkedlist) Print() {
	// This method Prints all the elements in the list
	current := l.Head
	for current.Next != nil {
		fmt.Printf("%v ", current.Val)
		current = current.Next
	}
	fmt.Printf("%v\n", current.Val)
}

func (l *DLinkedlist) append(data int) int {
	// By defination it should append the node to the end node
	new_node := Node{Val: data}

	if l.Head == nil {
		l.Head = &new_node
		new_node.Previous = nil
		return 1
	} else if l.Head.Next == nil {
		l.Head.Next = &new_node
		new_node.Previous = l.Head
	} else {
		current := l.Head
		for current.Next.Next != nil {
			current.Next = current.Next.Next
		}
		current.Next.Next = &new_node
		current.Next.Previous = current
		return 1
	}
	return 1

}

func (l *DLinkedlist) Pop() int {

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

	l := DLinkedlist{}
	l.append(5)

	// fmt.Printf("%v\n", l.Head)
	l.Print()
	l.append(25)
	l.Print()
	// l.Pop()
	// l.append(25)

	// l.Print()
}
