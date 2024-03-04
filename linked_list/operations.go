package linked_list

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

func (list *LinkedList) Sort(element int) {
	panic("NOT IMPLEMENTED")
}

func (list *LinkedList) Insert(element int, position int) {
	panic("NOT IMPLEMENTED")
}
func (list *LinkedList) Delete(element int, occurence int) bool {
	panic("NOT IMPLEMENTED")
}
