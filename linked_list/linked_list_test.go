package linked_list

import (
	"fmt"
	"testing"
)

func compareSlice(t *testing.T, slice1 []int, slice2 []int) bool {
	t.Helper()
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
func TestNewLinkedlist(t *testing.T) {

	list := NewLinkedList()
	testCases := []struct {
		name          string
		expectedState bool
	}{
		{name: fmt.Sprintf("New List shouldn't be of length %v", list.length), expectedState: list.length == 0},
		{name: "List head shouldn't be nil", expectedState: list.head != nil},
		{name: "List shouldn't be nil", expectedState: list != nil},
	}
	for _, tc := range testCases {
		if !tc.expectedState {
			t.Fatal(tc.name)
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

func TestToArray(t *testing.T) {
	array := []int{2, 3, 5, 3, 2}
	list := NewLinkedList()
	list.FromArray(array)
	if !compareSlice(t, array, list.ToArray()) {
		t.Fatalf("Both arrays should be equal but got : \noriginal array : %v \nresulting array : %v", array, list.ToArray())
	}
}

func TestSearch(t *testing.T) {
	array := []int{2, 3, 5, 3, 2, 1, 6, 9, 4, 232, 43, 2, 0}
	list := NewLinkedList()
	if list.Search(0) {
		t.Fatalf(`List should be emnpty`)
	}
	list.FromArray(array)
	if list.Search(22) {
		t.Fatalf("List shouldn't contain : %v", 22)
	}
	if !list.Search(232) {
		t.Fatalf("List should contain : %v, but not found", 232)
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
	if !compareSlice(t, array, arrayExpected) {
		t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
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
	if !compareSlice(t, array, arrayExpected) {
		t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
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
	if !compareSlice(t, modifiedArray, arrayExpected) {
		t.Fatalf("\nexpected ->> %v \nbut got ->> %v", arrayExpected, array)
	}

}
