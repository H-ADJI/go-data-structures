// Operations to perform multiple actions on the linked list data structure
package linked_list

import (
	"fmt"
)

// re-Initialize a linked list elements using a slice
func (list *LinkedList) FromArray(array []int) {
	list.head = new(listNode)
	list.length = 0
	currentNode := list.head
	var previous *listNode
	for _, el := range array {
		if previous != nil {
			currentNode = new(listNode)
			previous.next = currentNode
		}
		currentNode.data = el
		list.length++
		previous = currentNode
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
	iterator := list.Iterator()
	for iterator.HasNext() {
		current := iterator.GetNext()
		if current.data == element {
			return true
		}
	}
	return false
}

// TODO: Refactor swap and sort to act on listNode.data and not the listNode ptr<
func (list *LinkedList) Sort() {
	for i := 0; i < list.length; i++ {
		iterator := list.Iterator()
		for j := 0; j < list.length-i-1; j++ {
			current := iterator.GetNext()
			next := current.next
			if current.data > next.data {
				current.data, next.data = next.data, current.data
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
		if current.data == element {
			if current == list.head {
				list.head = &listNode{}
				list.length--
			} else {
				preceding.next = current.next
				list.length--
			}
			return true
		}
		preceding = current
	}
	return false
}
