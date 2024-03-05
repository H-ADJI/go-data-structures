// Linked lists structure implementation and some helper functions
package linked_list

import (
	"fmt"
	"strings"
)

const nodeTemplate = "Node(%v)"
const arrow = " --> "

type listNode struct {
	data int
	next *listNode
}

type LinkedList struct {
	head   *listNode
	length int
}

type LinkedListIterator struct {
	list    LinkedList
	current *listNode
}

func (node listNode) String() string {
	return fmt.Sprintf(nodeTemplate, node.data)
}
func NewLinkedList() *LinkedList {
	head := listNode{}
	return &LinkedList{head: &head, length: 0}
}
func (listIterator *LinkedListIterator) HasNext() bool {
	return listIterator.current.next != nil
}
func (listIterator *LinkedListIterator) next() {
	listIterator.current = listIterator.current.next
}

func (listIterator *LinkedListIterator) GetNext() *listNode {
	defer listIterator.next()
	return listIterator.current
}

func (list LinkedList) Iterator() *LinkedListIterator {
	return &LinkedListIterator{list: list, current: list.head}
}

func (list LinkedList) String() string {
	var stringBuidler strings.Builder
	iterator := list.Iterator()
	for iterator.HasNext() {
		node := iterator.GetNext()
		stringBuidler.WriteString(node.String())
		stringBuidler.WriteString(arrow)
	}
	return strings.TrimSuffix(stringBuidler.String(), arrow)
}
