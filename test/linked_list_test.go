package linked_list_test

import (
	"testing"

	linkedList "github.com/H-ADJI/go-data-structures/linked_list"
)

func TestNew(t *testing.T) {
	list := linkedList.New()
	if list == nil {
		t.Fatalf(`Wanted non-nil pointer for linked list, got %p pointer`, list)
	}
	if list == nil {
		t.Fatalf(`Wanted non-nil pointer for linked list, got %p pointer`, list)
	}
}
func TestSwap(t *testing.T) {
	array := []int{1, 2, 3, 4}
	list := linkedList.New()
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
func TestSearch(t *testing.T) {
	list := linkedList.New()
	if list.Search(0) {
		t.Fatalf(`List should be emnpty`)
	}
}
func TestSort(t *testing.T) {
	array := []int{1, 2, 3, 4}
	list := linkedList.New()
	list.FromArray(array)
	list.Sort()
	arrayExpected := []int{1, 2, 3, 4}
	for i := range arrayExpected {
		if array[i] != arrayExpected[i] {
			t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
		}

	}
	array = []int{1, 3, 2, 4}
	list = linkedList.New()
	list.FromArray(array)
	list.Sort()
	arrayExpected = []int{1, 2, 3, 4}
	for i := range arrayExpected {
		if array[i] != arrayExpected[i] {
			t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
		}

	}
}
func TestFromArray(t *testing.T) {
	array := []int{2, 3, 5, 3, 2, 1, 6, 9, 4, 232, 43, 2, 0}
	list := linkedList.New()
	list.FromArray(array)
	for _, element := range array {
		if !list.Search(element) {
			t.Fatalf("Element %v should be in the list", element)
		}
	}
}
