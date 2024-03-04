package linked_list

import (
	"fmt"
	"strings"
)

const nodeTemplate = "Node(%v)"
const arrow = " --> "

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

type LinkedListIterator struct {
	list    LinkedList
	current *Node
}

func (node Node) String() string {
	return fmt.Sprintf(nodeTemplate, node.data)
}
func New() *LinkedList {
	head := Node{}
	return &LinkedList{head: &head, length: 0}
}
func (listIterator *LinkedListIterator) hasNext() bool {
	return listIterator.current.next != nil
}
func (listIterator *LinkedListIterator) next() {
	listIterator.current = listIterator.current.next
}

func (listIterator *LinkedListIterator) getNext() *Node {
	defer listIterator.next()
	return listIterator.current
}

func (list LinkedList) iterator() *LinkedListIterator {
	return &LinkedListIterator{list: list, current: list.head}
}

func (list LinkedList) String() string {
	var stringBuidler strings.Builder
	iterator := list.iterator()
	for iterator.hasNext() {
		node := iterator.getNext()
		stringBuidler.WriteString(node.String())
		stringBuidler.WriteString(arrow)
	}
	return strings.TrimSuffix(stringBuidler.String(), arrow)
}
