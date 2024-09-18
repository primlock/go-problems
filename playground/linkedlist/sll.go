// Singly Linked-List Demo

package playground

import (
	"errors"
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Display() error {
	if l == nil {
		return errors.New("uninitialized list")
	}

	if l.head == nil {
		fmt.Println("empty list")
	}

	// Print out the nodes
	current := l.head
	for current != nil {

		if current.next != nil {
			fmt.Printf("%d->", current.val)
		} else {
			fmt.Printf("%d\n", current.val)
		}

		current = current.next
	}

	return nil
}

func (l *LinkedList) Add(val int) (*Node, error) {
	if l == nil {
		return &Node{
			val:  -1,
			next: nil,
		}, errors.New("uninitialized list")
	}

	node := &Node{
		val:  val,
		next: nil,
	}

	// There is no head
	if l.head == nil {
		l.head = node
		return node, nil
	}

	// There is a head
	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = node

	return node, nil
}

func (l *LinkedList) Remove(node *Node) error {
	if l == nil {
		return errors.New("uninitialized list")
	}

	if l.head == nil {
		return errors.New("empty list")
	}

	// Removing the head
	if node == l.head {
		l.head = l.head.next
		return nil
	}

	// Other than the head
	var prev *Node
	current := l.head

	for current != nil {
		if current == node {
			prev.next = current.next
			return nil
		}

		prev = current
		current = current.next
	}

	return errors.New("node not found")
}
