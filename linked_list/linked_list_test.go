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
	testCases := []struct {
		sourceArray   []int
		hasElement    int
		hasnotElement int
	}{
		{[]int{5, 5, 5, 5, 1, 5, 5}, 5, 4},
		{[]int{5, 5, 5, 5, 5, 5, 2}, 2, 1},
		{[]int{}, 4, 5},
	}

	list := NewLinkedList()
	for _, tc := range testCases {
		list.FromArray(tc.sourceArray)
		if !list.Search(tc.hasElement) {
			if list.length != 0 {
				t.Fatalf("List %v should contain %v", list, tc.hasElement)
			}
		}
		if list.Search(tc.hasnotElement) {
			t.Fatalf("List %v shouldn't contain %v", list, tc.hasElement)
		}
	}
}

func TestToArray(t *testing.T) {
	testCases := []struct {
		sourceArray []int
		outputArray []int
		equal       bool
	}{
		{[]int{1, 2, 3, 4, 1, 5, 5}, []int{1, 2, 3, 4, 1, 5, 5}, true},
		{[]int{5, 5, 5, 5, 1, 5, 5}, []int{}, false},
		{[]int{5, 5, 5, 5, 1, 5, 5}, []int{5, 1, 5, 5}, false},
		{[]int{}, []int{}, true},
		{[]int{0}, []int{0}, true},
		{[]int{}, []int{0}, false},
	}

	list := NewLinkedList()
	for _, tc := range testCases {
		list.FromArray(tc.sourceArray)
		if compareSlice(t, tc.outputArray, list.ToArray()) != tc.equal {
			if tc.equal {
				t.Fatalf("Expected to be different \nsource: %v, \noutput: %v", tc.sourceArray, tc.outputArray)
			} else {
				t.Fatalf("Expected to be equal \nsource: %v, \noutput: %v", tc.sourceArray, tc.outputArray)
			}
		}
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		element     int
		sourceArray []int
		in          bool
	}{
		{1, []int{1, 2, 3, 4, 1, 5, 5}, true},
		{0, []int{}, false},
		{0, []int{0}, true},
		{22, []int{5, 5}, false},
		{22, []int{22}, true},
		{42, []int{}, false},
		{360, []int{0}, false},
	}
	list := NewLinkedList()
	for _, tc := range testCases {
		list.FromArray(tc.sourceArray)
		if list.Search(tc.element) != tc.in {
			if tc.in {
				t.Fatalf("Expected to find the element: %v, in: %v", tc.element, tc.sourceArray)
			} else {
				t.Fatalf("Expected to not find the element: %v, in: %v", tc.element, tc.sourceArray)
			}
		}
	}

}
// func TestSwap(t *testing.T) {
// 	array := []int{1, 2, 3, 4, 5, 6}
// 	list := NewLinkedList()
// 	for i := 0; i < len(array)*2; i++ {
// 		list.FromArray(array)
// 		iterator := list.Iterator()
// 		for j := 0; j < i+1 && iterator.HasNext(); j++ {
// 			current := iterator.GetNext()
// 			if j == i-1 {
// 				preceding := current
// 			}
// 		}

// 	}
// }

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
