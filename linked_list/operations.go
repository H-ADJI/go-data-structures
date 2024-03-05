// Operations to perform multiple actions on the linked list data structure
package linked_list

import (
	"fmt"
)

func (list *LinkedList) FromArray(array []int) {
	currentNode := list.Head
	for _, el := range array {
		currentNode.data = el
		currentNode.next = new(Node)
		list.length++
		currentNode = currentNode.next
	}
}
func (list *LinkedList) ToArray() []int {
	iterator := list.Iterator()
	arr := []int{}
	for iterator.HasNext() {
		node := iterator.GetNext()
		arr = append(arr, node.data)
	}
	return arr
}
func (list *LinkedList) Search(element int) bool {
	if list.length == 0 {
		return false
	}
	current := list.Head
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}
	return false
}

func (node *Node) Swap(preceding *Node, following *Node) {
	if preceding == nil {
		node.next = following.next
		following.next = node
	} else {
		node.next = following.next
		following.next = node
		preceding.next = following
	}
}

func (list *LinkedList) Sort() {
	for i := 0; i < list.length; i++ {
		var preceding *Node = nil
		current := list.Head
		for j := 0; j < list.length-i-1; j++ {
			following := current.next
			if current.data > following.data {
				// swaping current and following node order
				current.Swap(preceding, following)
				preceding = following
			} else {
				preceding = current
				current = following
			}
		}

	}
}

func (list *LinkedList) Insert(element int, position int) error {
	node := &Node{data: element, next: nil}
	if position == 0 {
		if list.length == 0 {
			list.Head = node
		} else {
			node.next = list.Head
			list.Head = node
		}
	} else if position > 0 {
		if position > list.length {
			return fmt.Errorf("out of bound error, position %d is greater than list length %d", position, list.length)
		}
		var preceding *Node
		current := list.Head
		for i := 0; i < position+1; i++ {
			preceding = current
			current = current.next
		}
		preceding.next = node
		node.next = current
	}
	return nil

}
func (list *LinkedList) Delete(element int, occurence int) bool {
	panic("NOT IMPLEMENTED")
}
