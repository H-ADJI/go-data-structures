// Operations to perform multiple actions on the linked list data structure
package linked_list

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
		for j := 0; j < list.length-i; j++ {
			following := current.next
			if following != nil {
				if current.data > following.data {
					current.Swap(preceding, following)
					preceding = following
				} else {
					preceding = current
					current = following
				}
			}
		}

	}
}

func (list *LinkedList) Insert(element int, position int) {
	panic("NOT IMPLEMENTED")
}
func (list *LinkedList) Delete(element int, occurence int) bool {
	panic("NOT IMPLEMENTED")
}
