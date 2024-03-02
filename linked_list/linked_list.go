package linked_list

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head   *Node
	length int
}

func New() *LinkedList {
	head := Node{}
	return &LinkedList{head: &head, length: 0}
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
