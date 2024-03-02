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

func (node Node) String() string {
	return fmt.Sprintf(nodeTemplate, node.data)
}

type LinkedListIterator struct {
	list    LinkedList
	current *Node
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

type LinkedList struct {
	head   *Node
	length int
}

func New() *LinkedList {
	head := Node{}
	return &LinkedList{head: &head, length: 0}
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

func (list *LinkedList) FromArray(array []int) *LinkedList {
	currentNode := list.head
	for _, el := range array {
		currentNode.data = el
		currentNode.next = new(Node)
		list.length++
		currentNode = currentNode.next
	}
	return list
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
