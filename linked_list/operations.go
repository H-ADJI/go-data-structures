// Operations to perform multiple actions on the linked list data structure
package linked_list

import (
	"fmt"
)

func (list *LinkedList) FromArray(array []int) {
	currentNode := list.head
	for _, el := range array {
		currentNode.data = el
		currentNode.next = new(listNode)
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
	current := list.head
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}
	return false
}

func (node *listNode) Swap(preceding *listNode, following *listNode) {
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
		var preceding *listNode = nil
		current := list.head
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
	node := &listNode{data: element, next: nil}
	if position == 0 {
		if list.length == 0 {
			list.head = node
		} else {
			node.next = list.head
			list.head = node
		}
	} else if position > 0 {
		if position > list.length {
			return fmt.Errorf("out of bound error, position %d is greater than list length %d", position, list.length)
		}
		var preceding *listNode
		current := list.head
		for i := 0; i < position+1; i++ {
			preceding = current
			current = current.next
		}
		preceding.next = node
		node.next = current
	}
	return nil

}
func (list *LinkedList) Delete(element int) bool {
	var preceding *listNode
	iterator := list.Iterator()
	for iterator.HasNext() {
		current := iterator.GetNext()
		if list.head.data == element {
			list.head = list.head.next
			return true
		} else if current.data == element {
			preceding.next = current.next
			return true
		}
		preceding = current
	}
	return false
}
