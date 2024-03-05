package linked_list

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := NewLinkedList()
	if list == nil {
		t.Fatalf(`Wanted non-nil pointer for linked list, got %p pointer`, list)
	}
	if list == nil {
		t.Fatalf(`Wanted non-nil pointer for linked list, got %p pointer`, list)
	}
}
func TestSwap(t *testing.T) {
	array := []int{1, 2, 3, 4}
	list := NewLinkedList()
	list.FromArray(array)
	if list == nil {
		t.Fatalf(`Wanted non-nil pointer for linked list, got %p pointer`, list)
	}
	iterator := list.Iterator()
	if iterator.HasNext() {
		preceding := iterator.GetNext()
		current := iterator.GetNext()
		following := iterator.GetNext()
		current.Swap(preceding, following)
	}
	arrayExpected := []int{1, 3, 2, 4}
	array = list.ToArray()
	for i := range arrayExpected {
		if array[i] != arrayExpected[i] {
			t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
		}
	}
}
func TestFromArray(t *testing.T) {
	array := []int{2, 3, 5, 3, 2, 1, 6, 9, 4, 232, 43, 2, 0}
	list := NewLinkedList()
	list.FromArray(array)
	for _, element := range array {
		if !list.Search(element) {
			t.Fatalf("Element %v should be in the list", element)
		}
	}
}

func TestSearch(t *testing.T) {
	list := NewLinkedList()
	if list.Search(0) {
		t.Fatalf(`List should be emnpty`)
	}
}
func TestSort(t *testing.T) {
	array := []int{1, 2, 3, 4}
	list := NewLinkedList()
	list.FromArray(array)
	list.Sort()
	arrayExpected := []int{1, 2, 3, 4}
	for i := range arrayExpected {
		if array[i] != arrayExpected[i] {
			t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
		}

	}
	array = []int{1, 3, 2, 4}
	list = NewLinkedList()
	list.FromArray(array)
	list.Sort()
	arrayExpected = []int{1, 2, 3, 4}
	array = list.ToArray()
	for i := range arrayExpected {
		if array[i] != arrayExpected[i] {
			t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
		}

	}
}
func TestInsert(t *testing.T) {
	array := []int{2, 3, 5, 3, 2, 1, 6, 9, 4, 232, 43, 2, 0}
	list := NewLinkedList()
	list.FromArray(array)
	err := list.Insert(99, 88)
	if err == nil {
		t.Fatalf("The insert should return an error \n %v", list)
	}
	err = list.Insert(99, 0)
	if err != nil {
		t.Fatalf("element 99 should be the first element of the list but we got \n %v", list)
	}
	modifiedList := NewLinkedList()
	modifiedArray := []int{99, 2, 3, 5, 3, 2, 1, 6, 9, 4, 232, 43, 2, 0}
	modifiedList.FromArray(modifiedArray)
	err = modifiedList.Insert(99, 0)
	if err != nil {
		t.Fatalf("could not insert into list")
	}
	arrayExpected := list.ToArray()
	for i := range arrayExpected {
		if modifiedArray[i] != arrayExpected[i] {
			t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, modifiedArray)
		}

	}

}
